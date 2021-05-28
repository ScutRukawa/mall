package service

import (
	"context"
	"encoding/json"
	"online-mall-inventory/global"
	"online-mall-inventory/model"
	"online-mall-inventory/proto"

	"github.com/apache/rocketmq-client-go/v2/primitive"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Reback_inv(msg *primitive.MessageExt) error {
	//分布式锁
	orderMsgBody := model.OrderMsgBody{}
	err := json.Unmarshal(msg.Body, &orderMsgBody)
	if err != nil {
		zap.S().Error("messgae body parse error:", err) //todo
		return err
	}
	//查询库存扣减记录，然后归还
	invHistory := model.InventoryHistory{}
	ok, err := global.Engine.Where("order_id=?", orderMsgBody.OrderId).Get(&invHistory)
	zap.S().Info("查询扣减历史记录：", invHistory, ok)
	if err != nil || !ok {
		zap.S().Error("获取扣减历史失败：", err)
		return nil //库存未扣减，无需进行归还，确认消费
	}
	if invHistory.Status == model.Reback {
		return nil //库存已归还 ,无需重复归还，确认消费
	}
	orderIneDetial := model.OrderInvDetial{}
	err = json.Unmarshal(invHistory.OrderInvDetail, &orderIneDetial)
	if err != nil {
		return err
	}
	session := global.Engine.NewSession()
	session.Begin()
	for k, v := range orderIneDetial.GoodsNumsMap {
		sql := "update inventory set num=num+? where goods_id=? "
		result, err := session.Exec(sql, v.(float64), k)
		zap.S().Info("reback result :", result)
		//affected, _ := result.RowsAffected()
		if err != nil {
			zap.S().Info("归还库存失败：", err)
			session.Rollback()
			return err
		}
		sql = "update inventory_history set status=? where order_id=? "
		result, err = session.Exec(sql, int(model.Reback), orderMsgBody.OrderId)
		zap.S().Info("reback result :", result)
		//affected, _ := result.RowsAffected()
		if err != nil {
			zap.S().Info("写入归还记录失败：", err)
			session.Rollback()
			return err
		}
		zap.S().Info("归还订单:", orderMsgBody.OrderId)
	}
	if err = session.Commit(); err != nil {
		zap.S().Error("库存归还事务提交失败：", err)
		return err
	}
	return nil
}

type InventoryService struct {
}

func (i *InventoryService) SetInv(ctx context.Context, request *proto.GoodsInvInfo) (*proto.CommonRsp, error) {
	rsp := proto.CommonRsp{Code: proto.RetCode_SUCCESS}
	inventory := model.Inventory{}
	has, _ := global.Engine.Where("goods_id=?", request.GoodsId).Get(&inventory)
	if has {
		inventory.Num = request.Num
		_, err := global.Engine.Where("goods_id=?", request.GoodsId).Update(&inventory)
		if err != nil {
			rsp.Code = proto.RetCode_ERROR
			return &rsp, err
		}
		return &rsp, nil
	}
	inventory.GoodsId = request.GoodsId
	inventory.Num = request.Num
	_, err := global.Engine.Insert(&inventory)
	if err != nil {
		return &rsp, err
	}
	return &rsp, nil
}

func (i *InventoryService) GetInv(ctx context.Context, request *proto.GoodsInvInfo) (*proto.GoodsInvInfoRsp, error) {
	inventory := model.Inventory{}
	_, err := global.Engine.Where("goods_id=?", request.GoodsId).Get(&inventory)
	if err != nil {
		return nil, err
	}
	return &proto.GoodsInvInfoRsp{Num: inventory.Num}, nil
}

func (i *InventoryService) Sell(ctx context.Context, request *proto.SellInfo) (*proto.SellRsp, error) {
	zap.S().Info("recieve request: ", request.GoodsInfo)
	rsp := proto.SellRsp{Code: proto.RetCode_SUCCESS}
	invDetail := model.NewOrderInvDetial()
	session := global.Engine.NewSession()
	session.Begin()
	for _, item := range request.GoodsInfo {
		zap.S().Info("Sell item:", item)
		zap.S().Info("Sell item.goods_id:", item.GoodsId)

		// for {
		inventory := model.Inventory{}
		_, _ = session.Where("goods_id=?", item.GoodsId).Get(&inventory)
		if inventory.Num < item.Num {
			zap.S().Infof("inventory.Num:%d,item.Num:%d", inventory.Num, item.Num)
			session.Rollback()
			rsp.Code = proto.RetCode_inventory_insufficient
			return &rsp, status.Errorf(codes.ResourceExhausted, "库存不足")

		} else {
			sql := "update inventory set num=num-? where goods_id=? and version=? and num-?>=0"
			result, err := session.Exec(sql, item.Num, item.GoodsId, inventory.Version, item.Num)
			affected, _ := result.RowsAffected()
			if err != nil || affected == 0 {
				zap.S().Info("更新失败")
				session.Rollback()
				rsp.Code = proto.RetCode_ERROR
				return &rsp, status.Errorf(codes.Internal, "更新库存失败")
			}
			invDetail.GoodsNumsMap[item.GoodsId] = item.Num
			// }
		}
	}
	invDetailBlob, err := json.Marshal(invDetail)
	if err != nil {
		session.Rollback()
		return &rsp, status.Errorf(codes.Internal, "解析订单商品详情失败")
	}
	invHistory := model.InventoryHistory{
		OrderId:        request.OrderId,
		OrderInvDetail: invDetailBlob,
		Status:         model.OutStock,
	}
	affected, err := global.Engine.Insert(&invHistory)
	if err != nil || affected == 0 {
		session.Rollback()
		zap.S().Error("写入扣减历史记录失败:", err)
		return &rsp, status.Errorf(codes.Internal, "写入扣减历史记录失败")
	}
	err = session.Commit()
	zap.S().Info("transaction commit", err)
	return &rsp, nil
}

func (i *InventoryService) Reback(ctx context.Context, request *proto.SellInfo) (*proto.CommonRsp, error) {
	rsp := proto.CommonRsp{Code: proto.RetCode_SUCCESS}
	session := global.Engine.NewSession()
	session.Begin()
	for _, item := range request.GoodsInfo {
		inventory := model.Inventory{}
		_, err := session.Where("goods_id=?", item.GoodsId).Get(&inventory)
		if err != nil {
			session.Rollback()
		}
		inventory.Num += item.Num
		_, _ = session.Where("goods_id=?", item.GoodsId).Update(&inventory)
	}
	session.Commit()
	return &rsp, nil
}

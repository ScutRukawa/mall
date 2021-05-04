package service

import (
	"context"
	"errors"
	"online-mall-inventory/global"
	"online-mall-inventory/model"
	"online-mall-inventory/proto"

	"go.uber.org/zap"
)

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
	session := global.Engine.NewSession()
	session.Begin()
	for _, item := range request.GoodsInfo {
		zap.S().Info("Sell item:", item)
		// for {
		inventory := model.Inventory{}
		_, _ = session.Where("goods_id=?", item.GoodsId).Get(&inventory)
		if inventory.Num < item.Num {
			zap.S().Infof("inventory.Num:%d,item.Num:%d", inventory.Num, item.Num)
			session.Rollback()
			rsp.Code = proto.RetCode_inventory_insufficient
			return &rsp, errors.New("库存不足")
		} else {
			sql := "update inventory set num=num-? where goods_id=? and version=? and num-?>=0"
			result, err := session.Exec(sql, item.Num, item.GoodsId, inventory.Version, item.Num)
			affected, _ := result.RowsAffected()
			if err != nil || affected == 0 {
				zap.S().Info("更新失败")
				session.Rollback()
				rsp.Code = proto.RetCode_ERROR
				return &rsp, errors.New("更新失败")
			}
			// }
		}
	}
	err := session.Commit()
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

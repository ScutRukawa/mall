package goods

import (
	"context"
	"goodssrv/global"
	"goodssrv/model"
	"goodssrv/proto"

	"go.uber.org/zap"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Goods struct {
}

func (g *Goods) GoodsList(ctx context.Context, request *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	return nil, nil
}

//用户提交订单有多个商品，批量查询商品
func (g *Goods) BatchGetGoods(ctx context.Context, request *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
	zap.S().Info("BatchGetGoods recieve a request:", request)
	rsp := proto.GoodsListResponse{}
	goossInfoRsp := make([]*proto.GoodsInfoResponse, 0)
	for _, id := range request.Id {
		goods := model.Goods{}
		global.Engine.Where("id=?", id).Get(&goods)
		goodsInfo := proto.GoodsInfoResponse{
			Name:      goods.Name,
			Id:        int32(goods.Id),
			ShopPrice: float32(goods.ShopPrice),
		}
		goossInfoRsp = append(goossInfoRsp, &goodsInfo)
	}
	rsp.Data = goossInfoRsp
	rsp.Total = int32(len(goossInfoRsp))
	zap.S().Info("rsp.Data:", rsp.Data)
	return &rsp, nil
}
func (g *Goods) CreateGoods(ctx context.Context, request *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
	return nil, nil
}
func (g *Goods) DeleteGoods(ctx context.Context, request *proto.DeleteGoodsInfo) (*emptypb.Empty, error) {
	return nil, nil
}
func (g *Goods) UpdateGoods(ctx context.Context, request *proto.CreateGoodsInfo) (*emptypb.Empty, error) {
	return nil, nil
}
func (g *Goods) GetGoodsDetail(ctx context.Context, request *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
	return nil, nil
}

//商品分类
func (g *Goods) GetAllCategorysList(ctx context.Context, request *emptypb.Empty) (*proto.CategoryListResponse, error) {
	return nil, nil
}

//获取子分类
func (g *Goods) GetSubCategory(ctx context.Context, request *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	return nil, nil
}
func (g *Goods) CreateCategory(ctx context.Context, request *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	return nil, nil
}
func (g *Goods) DeleteCategory(ctx context.Context, request *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	return nil, nil
}
func (g *Goods) UpdateCategory(ctx context.Context, request *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	return nil, nil
}

//品牌和轮播图
func (g *Goods) BrandList(ctx context.Context, request *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	return nil, nil
}
func (g *Goods) CreateBrand(ctx context.Context, request *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	return nil, nil
}
func (g *Goods) DeleteBrand(ctx context.Context, request *proto.BrandRequest) (*emptypb.Empty, error) {
	return nil, nil
}
func (g *Goods) UpdateBrand(ctx context.Context, request *proto.BrandRequest) (*emptypb.Empty, error) {
	return nil, nil
}

//轮播图
func (g *Goods) BannerList(ctx context.Context, request *emptypb.Empty) (*proto.BannerListResponse, error) {
	return nil, nil
}
func (g *Goods) CreateBanner(ctx context.Context, request *proto.BannerRequest) (*proto.BannerResponse, error) {
	return nil, nil
}
func (g *Goods) DeleteBanner(ctx context.Context, request *proto.BannerRequest) (*emptypb.Empty, error) {
	return nil, nil
}
func (g *Goods) UpdateBanner(ctx context.Context, request *proto.BannerRequest) (*emptypb.Empty, error) {
	return nil, nil
}

//品牌分类
func (g *Goods) CategoryBrandList(ctx context.Context, request *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error) {
	return nil, nil
}

//通过category获取brands
func (g *Goods) GetCategoryBrandList(ctx context.Context, request *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {
	return nil, nil
}
func (g *Goods) CreateCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {
	return nil, nil
}
func (g *Goods) DeleteCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	return nil, nil
}
func (g *Goods) UpdateCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	return nil, nil
}

package logic

import (
	"context"
	"errors"
	"go_mall/apps/user/rpc/model"
	"google.golang.org/grpc/codes"
	"strings"

	"go_mall/apps/product/rpc/internal/svc"
	pb "go_mall/apps/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(in *pb.GetProductReq) (*pb.GetProductResp, error) {
	product, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		logx.Errorf("l.svcCtx.ProductModel.FindOne err:%#v", err)
		if errors.Is(err, model.ErrNotFound) {
			return &pb.GetProductResp{
				Code:    int64(codes.NotFound),
				Message: "商品不存在",
			}, nil
		}
		return &pb.GetProductResp{
			Code:    int64(codes.Unknown),
			Message: "系统错误，请联系管理员",
		}, nil
	}

	return &pb.GetProductResp{
		Code:      int64(codes.OK),
		Message:   "获取商品成功",
		Name:      product.Name,
		Subtitle:  product.Subtitle,
		Stock:     product.Stock,
		Price:     product.Price,
		Status:    product.Status,
		Detail:    product.Detail.String,
		Images:    strings.Split(product.Images.String, ","),
		Id:        int64(product.Id),
		CreatedAt: product.CreateTime.Format("2006-01-02 15:04:05"),
		UpdatedAt: product.UpdateTime.Format("2006-01-02 15:04:05"),
	}, nil
}

package product

import (
	"context"
	pb "go_mall/apps/product/rpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go_mall/apps/product/api/internal/svc"
	"go_mall/apps/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取产品线信息
func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductLogic) GetProduct(req *types.GetProductReq) (resp *types.Response, err error) {
	product, err := l.svcCtx.ProductRpc.GetProduct(l.ctx, &pb.GetProductReq{Id: req.Id})
	if err != nil {
		if status, ok := status.FromError(err); ok {
			return &types.Response{
				Code:    int64(status.Code()),
				Message: status.Message(),
			}, nil
		}
		return &types.Response{
			Code:    int64(codes.Internal),
			Message: "系统错误，请联系管理员",
		}, nil
	}
	return &types.Response{
		Code:    int64(product.GetCode()),
		Message: product.GetMessage(),
		Data:    product,
	}, nil
}

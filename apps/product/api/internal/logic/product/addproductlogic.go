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

type AddProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加产品信息
func NewAddProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLogic {
	return &AddProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProductLogic) AddProduct(req *types.AddProductReq) (resp *types.Response, err error) {
	logx.Debugf("添加产品信息%+v", req)
	pbResp, err := l.svcCtx.ProductRpc.AddProduct(l.ctx, &pb.AddProductReq{
		CateId:   req.CateId,
		Name:     req.Name,
		Price:    req.Price,
		Status:   req.Status,
		Stock:    req.Stock,
		Subtitle: req.Subtitle,
		Detail:   req.Detail,
		Images:   req.Images,
	})

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
		Code:    int64(pbResp.GetCode()),
		Message: pbResp.GetMessage(),
	}, nil
}

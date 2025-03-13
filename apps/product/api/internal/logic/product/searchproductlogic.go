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

type SearchProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 搜索产品信息
func NewSearchProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductLogic {
	return &SearchProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchProductLogic) SearchProduct(req *types.SearchProductReq) (resp *types.Response, err error) {
	product, err := l.svcCtx.ProductRpc.SearchProduct(l.ctx, &pb.SearchProductReq{
		Keyword: req.Keyword,
		Page:    req.Page,
		Size:    req.Size,
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
		Code:    int64(codes.OK),
		Message: "查询成功",
		Data:    product,
	}, nil
}

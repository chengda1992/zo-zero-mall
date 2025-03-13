package logic

import (
	"context"

	"go_mall/apps/product/rpc/internal/svc"
	"go_mall/apps/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductLogic {
	return &SearchProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchProductLogic) SearchProduct(in *__.SearchProductReq) (*__.SearchProductResp, error) {
	// todo: add your logic here and delete this line

	return &__.SearchProductResp{}, nil
}

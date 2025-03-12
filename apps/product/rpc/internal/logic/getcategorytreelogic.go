package logic

import (
	"context"

	"go_mall/apps/product/rpc/internal/svc"
	"go_mall/apps/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryTreeLogic {
	return &GetCategoryTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryTreeLogic) GetCategoryTree(in *__.GetCategoryTreeReq) (*__.GetCategoryTreeResp, error) {
	// todo: add your logic here and delete this line

	return &__.GetCategoryTreeResp{}, nil
}

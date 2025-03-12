package logic

import (
	"context"

	"go_mall/apps/product/rpc/internal/svc"
	"go_mall/apps/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(in *__.DeleteCategoryReq) (*__.DeleteCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &__.DeleteCategoryResp{}, nil
}

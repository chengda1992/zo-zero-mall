package logic

import (
	"context"

	"go_mall/apps/product/rpc/internal/svc"
	"go_mall/apps/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *__.UpdateCategoryReq) (*__.UpdateCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &__.UpdateCategoryResp{}, nil
}

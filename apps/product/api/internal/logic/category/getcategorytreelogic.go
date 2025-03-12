package category

import (
	"context"

	"go_mall/apps/product/api/internal/svc"
	"go_mall/apps/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分类树
func NewGetCategoryTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryTreeLogic {
	return &GetCategoryTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryTreeLogic) GetCategoryTree() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

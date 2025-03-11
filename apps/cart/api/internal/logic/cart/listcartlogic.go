package cart

import (
	"context"

	"go_mall/apps/cart/api/internal/svc"
	"go_mall/apps/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取购物车列表
func NewListCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCartLogic {
	return &ListCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCartLogic) ListCart(req *types.ListReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

package cart

import (
	"context"

	"go_mall/apps/cart/api/internal/svc"
	"go_mall/apps/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新购物车
func NewUpdateCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartLogic {
	return &UpdateCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCartLogic) UpdateCart(req *types.UpdateReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

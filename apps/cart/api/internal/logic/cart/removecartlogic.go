package cart

import (
	"context"

	"go_mall/apps/cart/api/internal/svc"
	"go_mall/apps/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除购物车
func NewRemoveCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveCartLogic {
	return &RemoveCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveCartLogic) RemoveCart(req *types.RemoveReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

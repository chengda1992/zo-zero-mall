package cart

import (
	"context"
	"errors"
	"go_mall/apps/cart/rpc/cart"

	"go_mall/apps/cart/api/internal/svc"
	"go_mall/apps/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加购物车
func NewAddCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartLogic {
	return &AddCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCartLogic) AddCart(req *types.AddReq) (resp *types.Response, err error) {
	add, err := l.svcCtx.CartRpc.Add(l.ctx, &cart.AddReq{
		Userid:   req.UserID,
		Proid:    req.ProID,
		Quantity: req.Quantity,
		Checked:  req.Checked,
	})
	if err != nil {
		return nil, errors.New("添加购物车失败")
	}
	if add.GetCode() == 0 {
		return nil, errors.New("添加购物车失败")
	}

	return &types.Response{
		Code:    add.GetCode(),
		Message: add.GetMessage(),
	}, nil
}

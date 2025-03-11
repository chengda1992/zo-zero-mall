package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go_mall/apps/cart/api/internal/config"
	"go_mall/apps/cart/rpc/cart"
)

type ServiceContext struct {
	Config  config.Config
	CartRpc cart.Cart
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		CartRpc: cart.NewCart(zrpc.MustNewClient(c.CartRpc)),
	}
}

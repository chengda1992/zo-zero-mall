package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go_mall/apps/product/api/internal/config"
	"go_mall/apps/product/rpc/product"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}

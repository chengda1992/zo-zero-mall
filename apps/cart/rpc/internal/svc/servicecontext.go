package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_mall/apps/cart/rpc/internal/config"
	"go_mall/apps/cart/rpc/model"
)

type ServiceContext struct {
	Config    config.Config
	CartModel model.CartModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := sqlx.NewMysql(c.DB)
	return &ServiceContext{
		Config:    c,
		CartModel: model.NewCartModel(mysql),
	}
}

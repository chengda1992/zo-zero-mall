package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_mall/apps/product/rpc/internal/config"
	"go_mall/apps/product/rpc/model"
)

type ServiceContext struct {
	Config config.Config

	CategoryModel model.CategoryModel
	ProductModel  model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := sqlx.NewMysql(c.DB)
	return &ServiceContext{
		Config:        c,
		CategoryModel: model.NewCategoryModel(mysql, c.CacheRedis),
		ProductModel:  model.NewProductModel(mysql),
	}
}

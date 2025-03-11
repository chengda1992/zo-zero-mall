package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_mall/apps/user/rpc/internal/config"
	"go_mall/apps/user/rpc/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := sqlx.NewMysql(c.DB)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(mysql),
	}
}

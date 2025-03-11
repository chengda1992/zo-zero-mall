package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"go_mall/apps/user/api/internal/config"
	"go_mall/apps/user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := zrpc.MustNewClient(c.UserRpc)
	logx.Infof("Consul resolved address: %s", client.Conn().Target())
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(client),
	}
}

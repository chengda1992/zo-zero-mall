package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf

	Consul consul.Conf

	DB string

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}

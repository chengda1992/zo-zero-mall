package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf

	Consul consul.Conf

	CacheRedis cache.CacheConf

	DB string

	Elasticsearch struct {
		Address string `json:"Address"`
	}
}

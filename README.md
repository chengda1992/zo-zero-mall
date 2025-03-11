## go-mall 商城项目

go-zero微服务实战项目
分为5个模块，分别是cart order product user pay

## 1 在项目的根目录下
```shell
go mod init go_mall
cd apps
# 生成cart api
goctl api go -api cart/api/cart.api -dir cart/api
#生成cart rpc  和cart.proto同级目录
cd cart/rpc/pb
goctl rpc protoc cart.proto --go_out=../pb --go-grpc_out=../pb --zrpc_out=../../rpc
#生成cart model
cd ..
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/mall-user" -table="user" -dir="./model" --style=goZero

```
## 2 注册中心改为consul
```shell
# rpc yaml
Consul:
    Host: 127.0.0.1:8500
    Key: user_rpc

#Config.go config配置文件添加 Consul consul.Conf
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	Consul consul.Conf
main.go
	#注册服务到consul
	consul.RegisterService(c.ListenOn, c.Consul)
#api yaml
UserRpc:
    Target: consul://127.0.0.1:8500/user_rpc?wait=15s
main.go
		_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul" # 这个很重要，否则会报错 告诉用consul

```



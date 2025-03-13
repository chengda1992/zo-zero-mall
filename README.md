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
goctl rpc protoc product.proto --go_out=. --go-grpc_out=. --zrpc_out=../
#生成cart model
cd ..
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/mall-product" -table="product" -dir="./model" --style=goZero
# -c 开启缓存
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/mall-user" -table="user" -dir="./model" --style=goZero -c

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

## 3 增加es


 Elasticsearch
在 go-zero 的配置文件中增加 Elasticsearch 的配置项。

配置文件示例（config.yaml）
yaml
Elasticsearch:
  Address: "http://localhost:9200" # ES 地址
配置结构体（config/config.go）
go
type Config struct {
    Elasticsearch struct {
        Address string `json:"Address"`
    }
    // 其他配置...
}
3. 初始化 Elasticsearch 客户端
在 svc.ServiceContext 中初始化 Elasticsearch 客户端，以便在逻辑层中使用。

代码示例（svc/servicecontext.go）
go
import (
    "github.com/elastic/go-elasticsearch/v8"
    "github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
    Config      config.Config
    ProductModel model.ProductModel
    EsClient    *elasticsearch.Client // ES 客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
    // 初始化 Elasticsearch 客户端
    esClient, err := elasticsearch.NewClient(elasticsearch.Config{
        Addresses: []string{c.Elasticsearch.Address},
    })
    if err != nil {
        logx.Errorf("Failed to create Elasticsearch client: %v", err)
    }

    return &ServiceContext{
        Config:      c,
        ProductModel: model.NewProductModel(sqlx.NewMysql(c.Mysql.DataSource)),
        EsClient:    esClient,
    }
}




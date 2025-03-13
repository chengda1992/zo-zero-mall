package svc

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_mall/apps/product/rpc/internal/config"
	"go_mall/apps/product/rpc/model"
)

type ServiceContext struct {
	Config config.Config

	CategoryModel model.CategoryModel
	ProductModel  model.ProductModel

	EsClient *elasticsearch.Client // ES 客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := sqlx.NewMysql(c.DB)

	// 初始化 Elasticsearch 客户端
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{c.Elasticsearch.Address},
	})
	if err != nil {
		logx.Errorf("Failed to create Elasticsearch client: %v", err)
	}

	return &ServiceContext{
		Config:        c,
		CategoryModel: model.NewCategoryModel(mysql, c.CacheRedis),
		ProductModel:  model.NewProductModel(mysql),
		EsClient:      esClient,
	}
}

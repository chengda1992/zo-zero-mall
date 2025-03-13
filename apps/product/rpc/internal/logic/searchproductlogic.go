package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"go_mall/apps/product/rpc/internal/svc"
	pb "go_mall/apps/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductLogic {
	return &SearchProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchProductLogic) SearchProduct(in *pb.SearchProductReq) (*pb.SearchProductResp, error) {
	// 构建 Elasticsearch 查询
	type WildcardQuery struct {
		Value string `json:"value"`
	}

	type Query struct {
		Wildcard map[string]WildcardQuery `json:"wildcard"`
	}

	type SearchQuery struct {
		Query Query `json:"query"`
		From  int   `json:"from"`
		Size  int   `json:"size"`
	}

	query := SearchQuery{
		Query: Query{
			Wildcard: map[string]WildcardQuery{
				"name": {
					Value: "*" + in.Keyword + "*",
				},
			},
		},
		From: int((in.Page - 1) * in.Size),
		Size: int(in.Size),
	}

	queryJSON, _ := json.Marshal(query)

	// 执行查询
	res, err := l.svcCtx.EsClient.Search(
		l.svcCtx.EsClient.Search.WithIndex("product"),
		l.svcCtx.EsClient.Search.WithBody(bytes.NewReader(queryJSON)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// 解析查询结果
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		logx.Errorf("json.NewDecoder.Decode err: %#V", err)
		return nil, err
	}

	// 构建响应
	var products []*pb.ProductInfo
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})

		var images []string
		switch v := source["images"].(type) {
		case string:
			images = strings.Split(v, ",")
		case []interface{}:
			for _, img := range v {
				if str, ok := img.(string); ok {
					images = append(images, str)
				}
			}
		default:
			// 处理其他类型或记录日志
			l.Logger.Errorf("unexpected type for images: %T", v)
		}

		products = append(products, &pb.ProductInfo{
			Id:       int64(source["id"].(float64)),
			CateId:   int64(source["cate_id"].(float64)),
			Name:     source["name"].(string),
			Subtitle: source["subtitle"].(string),
			Detail:   source["detail"].(string),
			Price:    source["price"].(float64),
			Stock:    int64(source["stock"].(float64)),
			Status:   int64(source["status"].(float64)),
			Images:   images,
		})
	}
	total := int64(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	hasNext := (in.Page * in.Size) < total
	return &pb.SearchProductResp{
		Products: products,
		Total:    total,
		HaxNext:  hasNext,
	}, nil

}

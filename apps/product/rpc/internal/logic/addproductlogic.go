package logic

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"go_mall/apps/product/rpc/internal/svc"
	"go_mall/apps/product/rpc/model"
	pb "go_mall/apps/product/rpc/pb"
	"google.golang.org/grpc/codes"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLogic {
	return &AddProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProductLogic) AddProduct(in *pb.AddProductReq) (*pb.AddProductResp, error) {
	logx.Debugf("添加的入参%+v", in)
	result, err := l.svcCtx.ProductModel.Insert(l.ctx, &model.Product{
		Cateid:   uint64(in.CateId),
		Detail:   sql.NullString{String: in.Detail, Valid: true},
		Images:   sql.NullString{String: strings.Join(in.Images, ","), Valid: true},
		Name:     in.Name,
		Price:    in.Price,
		Status:   int64(in.Status),
		Stock:    int64(in.Stock),
		Subtitle: in.Subtitle,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.ProductModel.Insert err:%#v", err)
		return &pb.AddProductResp{
			Code:    int64(codes.Unknown),
			Message: "系统错误，请联系管理员",
		}, nil

	}
	id, err := result.LastInsertId()
	if err != nil {
		logx.Errorf("l.svcCtx.ProductModel.Insert err:%#v", err)
		return &pb.AddProductResp{
			Code:    int64(codes.Unknown),
			Message: "系统错误，请联系管理员",
		}, nil
	}
	// 2. 将产品数据同步到 Elasticsearch
	esProduct := map[string]interface{}{
		"id":          id,
		"cate_id":     in.CateId,
		"name":        in.Name,
		"subtitle":    in.Subtitle,
		"detail":      in.Detail,
		"price":       in.Price,
		"stock":       in.Stock,
		"status":      in.Status,
		"images":      in.Images,
		"create_time": time.Now().Format(time.RFC3339),
		"update_time": time.Now().Format(time.RFC3339),
	}
	jsonData, err := json.Marshal(esProduct)
	if err != nil {
		log.Fatalf("Failed to marshal map: %v", err)
	}
	_, err = l.svcCtx.EsClient.Index("product", bytes.NewReader(jsonData),
		l.svcCtx.EsClient.Index.WithDocumentID(strconv.FormatInt(id, 10)))
	if err != nil {
		logx.Errorf("Failed to sync product to Elasticsearch: %v", err)
	}

	return &pb.AddProductResp{
			Code:    int64(codes.OK),
			Message: "添加成功"},
		nil
}

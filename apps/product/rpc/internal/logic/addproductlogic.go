package logic

import (
	"context"
	"database/sql"
	"go_mall/apps/product/rpc/internal/svc"
	"go_mall/apps/product/rpc/model"
	pb "go_mall/apps/product/rpc/pb"
	"google.golang.org/grpc/codes"
	"strings"

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
	_, err := l.svcCtx.ProductModel.Insert(l.ctx, &model.Product{
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

	return &pb.AddProductResp{
			Code:    int64(codes.OK),
			Message: "添加成功"},
		nil
}

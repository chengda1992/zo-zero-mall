package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"strconv"

	"go_mall/apps/product/rpc/internal/svc"
	pb "go_mall/apps/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryByIdLogic {
	return &GetCategoryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分类
func (l *GetCategoryByIdLogic) GetCategoryById(in *pb.GetCategoryReq) (*pb.GetCategoryResp, error) {
	category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return &pb.GetCategoryResp{
				Code:    int64(codes.NotFound),
				Message: "分类不存在",
			}, nil
		}

		logx.Errorf("l.svcCtx.CategoryModel.FindOne err: %#V", err)
		return &pb.GetCategoryResp{
			Code:    int64(codes.Unknown),
			Message: "系统错误，请联系管理员",
		}, nil
	}

	return &pb.GetCategoryResp{
		Code:    int64(codes.OK),
		Message: "success",
		Category: &pb.CategoryInfo{
			Id:         int64(category.Id),
			Name:       category.Name,
			Parentid:   strconv.FormatInt(category.Parentid, 10),
			Status:     int64(category.Status),
			CreateTime: category.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: category.UpdateTime.Format("2006-01-02 15:04:05"),
		},
	}, nil
}

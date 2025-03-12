package logic

import (
	"context"
	"go_mall/apps/product/rpc/internal/svc"
	"go_mall/apps/product/rpc/model"
	pb "go_mall/apps/product/rpc/pb"
	"google.golang.org/grpc/codes"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCategoryLogic) AddCategory(in *pb.AddCategoryReq) (*pb.AddCategoryResp, error) {
	if len(in.Name) == 0 {
		return &pb.AddCategoryResp{
			Code:    int64(codes.InvalidArgument),
			Message: "品类必传",
		}, nil
	}

	//查询当前的品类是否已经存在
	oldCategory, err := l.svcCtx.CategoryModel.FindOneByNameAndParentId(l.ctx, in.Name, in.ParentId)
	if err != nil && err != model.ErrNotFound {
		logx.Errorf("l.svcCtx.CategoryModel.FindOneByNameAndParentId err: %#V", err)
		return &pb.AddCategoryResp{
			Code:    int64(codes.Unknown),
			Message: "系统错误，请联系管理员",
		}, nil
	}
	if oldCategory != nil {
		return &pb.AddCategoryResp{
			Code:    int64(codes.AlreadyExists),
			Message: "该品类已经存在",
		}, nil
	}

	//添加品类到数据库
	_, err = l.svcCtx.CategoryModel.Insert(l.ctx, &model.Category{
		Name:     in.Name,
		Parentid: in.ParentId,
		Status:   1,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.CategoryModel.Insert err: %#V", err)
		return &pb.AddCategoryResp{
			Code:    int64(codes.Unknown),
			Message: "系统错误，请联系管理员",
		}, nil
	}

	return &pb.AddCategoryResp{
		Code:    int64(codes.OK),
		Message: "添加品类成功",
	}, nil
}

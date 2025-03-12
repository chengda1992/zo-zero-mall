package category

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go_mall/apps/product/api/internal/svc"
	"go_mall/apps/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	pb "go_mall/apps/product/rpc/pb"
)

const (
	// 父级分类
	PARENT_LEVEL = 0
)

type AddCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加分类
func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCategoryLogic) AddCategory(req *types.AddCategoryReq) (resp *types.Response, err error) {
	//默认父级分类
	//如果父级分类为0，则默认为父级分类
	if req.ParentId == 0 {
		req.ParentId = PARENT_LEVEL
	}

	category, err := l.svcCtx.ProductRpc.AddCategory(l.ctx, &pb.AddCategoryReq{
		ParentId: req.ParentId,
		Name:     req.Name,
	})

	if err != nil {
		logx.Errorf("l.svcCtx.ProductRpc.AddCategory err:%#v", err)
		if status, ok := status.FromError(err); ok {
			return &types.Response{
				Code:    int64(status.Code()),
				Message: status.Message(),
			}, nil
		}
		return &types.Response{
			Code:    int64(codes.Internal),
			Message: "系统错误，请联系管理员",
		}, nil
	}

	return &types.Response{
		Code:    int64(category.GetCode()),
		Message: category.Message,
	}, nil
}

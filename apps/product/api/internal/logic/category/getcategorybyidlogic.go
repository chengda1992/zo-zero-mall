package category

import (
	"context"
	pb "go_mall/apps/product/rpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go_mall/apps/product/api/internal/svc"
	"go_mall/apps/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分类详情
func NewGetCategoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryByIdLogic {
	return &GetCategoryByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryByIdLogic) GetCategoryById(req *types.GetCategoryByIdReq) (resp *types.Response, err error) {
	if req.Id == 0 {
		return &types.Response{
			Code:    int64(codes.InvalidArgument),
			Message: "id必传",
		}, nil
	}

	categoryResp, err := l.svcCtx.ProductRpc.GetCategoryById(l.ctx, &pb.GetCategoryReq{
		Id: req.Id,
	})

	if err != nil {
		if code, ok := status.FromError(err); ok {
			return &types.Response{
				Code:    int64(code.Code()),
				Message: code.Message(),
			}, nil
		}
		return &types.Response{
			Code:    int64(codes.Internal),
			Message: "系统错误，请联系管理员",
		}, nil
	}

	return &types.Response{
		Code:    int64(categoryResp.GetCode()),
		Message: categoryResp.GetMessage(),
		Data:    categoryResp.GetCategory(),
	}, nil
}

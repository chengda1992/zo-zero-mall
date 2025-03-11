package user

import (
	"context"
	"go_mall/apps/user/rpc/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go_mall/apps/user/api/internal/svc"
	"go_mall/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.InfoReq) (resp *types.Response, err error) {
	info, err := l.svcCtx.UserRpc.Info(l.ctx, &user.InfoReq{Id: req.ID})
	if err != nil {
		// 将gRPC错误转换为HTTP错误
		if st, ok := status.FromError(err); ok {
			return &types.Response{
				Code:    int64(st.Code()),
				Message: st.Message(),
			}, nil
		}
		return &types.Response{
			Code:    int64(codes.Internal),
			Message: "系统错误，请联系管理员",
		}, nil
	}
	return &types.Response{
		Code:    info.GetCode(),
		Message: info.GetMessage(),
		Data:    info.GetData(),
	}, nil
}

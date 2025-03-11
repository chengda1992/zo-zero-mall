package user

import (
	"context"
	"go_mall/apps/user/api/internal/svc"
	"go_mall/apps/user/api/internal/types"
	"go_mall/apps/user/rpc/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 注册用户
func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.Response, err error) {
	register, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Question: req.Question,
		Answer:   req.Answer,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.UserRpc.Register err: %#V", err)
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
		Code:    register.GetCode(),
		Message: register.GetMessage(),
	}, nil
}

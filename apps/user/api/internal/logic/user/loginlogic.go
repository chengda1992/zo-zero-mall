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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 用户登录
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.Response, err error) {

	//1.查询当前用户名是否存在
	user, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
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
		Code:    user.GetCode(),
		Message: user.GetMessage(),
		Data:    user.GetToken(),
	}, nil
}

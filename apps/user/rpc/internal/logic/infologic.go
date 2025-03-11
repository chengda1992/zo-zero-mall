package logic

import (
	"context"
	"errors"
	"go_mall/apps/user/rpc/model"
	"google.golang.org/grpc/codes"

	"go_mall/apps/user/rpc/internal/svc"
	pb "go_mall/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *pb.InfoReq) (*pb.InfoResp, error) {
	if in.Id == 0 {
		return &pb.InfoResp{
			Code:    int64(codes.InvalidArgument),
			Message: "id不能为空",
		}, nil
	}
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		logx.Errorf("l.svcCtx.UserModel.FindOne err: %#V", err)
		if errors.Is(err, model.ErrNotFound) {
			return &pb.InfoResp{
				Code:    int64(codes.NotFound),
				Message: "用户不存在",
			}, nil
		}
		return &pb.InfoResp{
			Code:    int64(codes.Internal),
			Message: "系统错误",
		}, nil
	}

	return &pb.InfoResp{
		Code:    int64(codes.OK),
		Message: "success",
		Data: &pb.UserInfo{
			Id:         int64(user.Id),
			Username:   user.Username,
			Phone:      user.Phone,
			CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
		},
	}, nil
}

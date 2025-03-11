package logic

import (
	"context"
	"go_mall/apps/user/rpc/internal/svc"
	"go_mall/apps/user/rpc/model"
	pb "go_mall/apps/user/rpc/pb"
	"go_mall/pkg/md5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	//1.查询当前手机号是否已经存在
	phone := in.GetPhone()
	user, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, phone)
	if err != nil && err != model.ErrNotFound {
		logx.Errorf("l.svcCtx.UserModel.FindOneByPhone err: %#V", err)
		return nil, status.Error(codes.Internal, "系统错误")
	}
	if user != nil {
		return nil, status.Error(codes.AlreadyExists, "手机号已经存在")
	}

	//2.对当前密码进行加密
	md5Pwd := md5.Sum(in.GetPassword())
	user = &model.User{
		Username: in.GetUsername(),
		Password: md5Pwd,
		Phone:    in.GetPhone(),
		Question: in.GetQuestion(),
		Answer:   in.GetAnswer(),
	}
	insert, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		logx.Errorf("l.svcCtx.UserModel.Insert err: %#V", err)
		return nil, status.Error(codes.Internal, "系统错误")
	}
	if count, err := insert.RowsAffected(); err != nil && count != 1 {
		return nil, status.Error(codes.Internal, "系统错误")
	}
	return &pb.RegisterResp{
		Code:    int64(codes.OK),
		Message: "注册成功",
	}, nil
}

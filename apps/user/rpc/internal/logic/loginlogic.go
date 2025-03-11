package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"go_mall/apps/user/rpc/internal/svc"
	"go_mall/apps/user/rpc/model"
	"go_mall/apps/user/rpc/pb/pb"
	"go_mall/pkg/md5"
	"google.golang.org/grpc/codes"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	//判断输入的用户名和密码是否为空
	if in.Username == "" || in.Password == "" {
		return &pb.LoginResp{
			Code:    int64(codes.InvalidArgument),
			Message: "用户名或密码不能为空",
		}, nil
	}
	//通过用户名查询用户是否存在
	user, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, in.Username)
	if err != nil && err != model.ErrNotFound {
		logx.Errorf("l.svcCtx.UserModel.FindOneByUserName err: %#V", err)
		return &pb.LoginResp{
			Code:    int64(codes.Internal),
			Message: "系统错误",
		}, nil
	}
	if user == nil {
		return &pb.LoginResp{
			Code:    int64(codes.NotFound),
			Message: "用户不存在",
		}, nil
	}

	//将当前的登录密码进行加密
	md5Pwd := md5.Sum(in.GetPassword())
	if user.Password != md5Pwd {
		return &pb.LoginResp{
			Code:    int64(codes.PermissionDenied),
			Message: "密码错误",
		}, nil
	}
	//登录成功
	secret := l.svcCtx.Config.JwtAuth.AccessSecret
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	unix := time.Now().Unix()
	token, err := getJwtToken(secret, unix, accessExpire, user)
	if err != nil {
		logx.Errorf("getJwtToken err: %#V", err)
		return &pb.LoginResp{
			Code:    int64(codes.Internal),
			Message: "系统错误",
		}, nil
	}
	return &pb.LoginResp{
		Code:    int64(codes.OK),
		Message: "登录成功",
		Data: &pb.UserInfo{
			Id:         int64(user.Id),
			Username:   user.Username,
			Phone:      user.Phone,
			CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
		},
		Token: token,
	}, nil
}

// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func getJwtToken(secretKey string, iat, seconds int64, payload interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

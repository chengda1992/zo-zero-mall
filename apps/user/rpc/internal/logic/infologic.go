package logic

import (
	"context"

	"go_mall/apps/user/rpc/internal/svc"
	"go_mall/apps/user/rpc/pb/pb"

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
	// todo: add your logic here and delete this line

	return &pb.InfoResp{}, nil
}

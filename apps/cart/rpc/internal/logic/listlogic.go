package logic

import (
	"context"

	"go_mall/apps/cart/rpc/internal/svc"
	pb "go_mall/apps/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *pb.ListReq) (*pb.ListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ListResp{}, nil
}

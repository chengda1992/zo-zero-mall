package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_mall/apps/cart/rpc/internal/svc"
	pb "go_mall/apps/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *pb.AddReq) (*pb.AddResp, error) {
	err := l.svcCtx.CartModel.Transact(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		logx.Info("事务处理")
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddResp{
		Code:    200,
		Message: "OK",
	}, nil
}

package logic

import (
	"context"

	"go_mall/apps/cart/rpc/internal/svc"
	"go_mall/apps/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *__.RemoveReq) (*__.RemoveResp, error) {
	// todo: add your logic here and delete this line

	return &__.RemoveResp{}, nil
}

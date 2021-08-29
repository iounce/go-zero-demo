package logic

import (
	"context"

	"errors"
	"push/internal/svc"
	"push/push"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubscribeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubscribeLogic {
	return &SubscribeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubscribeLogic) Subscribe(in *push.MsgReq) (*push.MsgRsp, error) {
	if len(in.Title) == 0 || len(in.Content) == 0 {
		return nil, errors.New("Invalid parameter")
	}

	println(in.Title, in.Content)

	return &push.MsgRsp{
		Code: 0,
		Msg:  "OK",
	}, nil
}

// Code generated by goctl. DO NOT EDIT!
// Source: msg.proto

package server

import (
	"context"

	"push/internal/logic"
	"push/internal/svc"
	"push/push"
)

type PushServer struct {
	svcCtx *svc.ServiceContext
}

func NewPushServer(svcCtx *svc.ServiceContext) *PushServer {
	return &PushServer{
		svcCtx: svcCtx,
	}
}

func (s *PushServer) Subscribe(ctx context.Context, in *push.MsgReq) (*push.MsgRsp, error) {
	l := logic.NewSubscribeLogic(ctx, s.svcCtx)
	return l.Subscribe(in)
}

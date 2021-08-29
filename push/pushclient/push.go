// Code generated by goctl. DO NOT EDIT!
// Source: msg.proto

//go:generate mockgen -destination ./push_mock.go -package pushclient -source $GOFILE

package pushclient

import (
	"context"

	"rpc/push"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	MsgReq = push.MsgReq
	MsgRsp = push.MsgRsp

	Push interface {
		Subscribe(ctx context.Context, in *MsgReq) (*MsgRsp, error)
	}

	defaultPush struct {
		cli zrpc.Client
	}
)

func NewPush(cli zrpc.Client) Push {
	return &defaultPush{
		cli: cli,
	}
}

func (m *defaultPush) Subscribe(ctx context.Context, in *MsgReq) (*MsgRsp, error) {
	client := push.NewPushClient(m.cli.Conn())
	return client.Subscribe(ctx, in)
}
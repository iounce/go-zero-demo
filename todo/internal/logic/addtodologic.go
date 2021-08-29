package logic

import (
	"context"
	"rpc/pushclient"

	"abc.com/todo/v1/internal/svc"
	"abc.com/todo/v1/internal/types"
	"abc.com/todo/v1/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddTodoLogic {
	return AddTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTodoLogic) AddTodo(req types.TodoRequest) (*types.TodoResponse, error) {
	data := model.Todo{Title: req.Title, Content: req.Content}
	result, err := l.svcCtx.TodoModel.Insert(data)
	if err != nil {
		return nil, err
	}

	var id int

	lastId, err := result.LastInsertId()
	if err == nil {
		id = int(lastId)
	}

	rpcReq := pushclient.MsgReq{Id: lastId, Title: req.Title, Content: req.Content}

	rpcRsp, err := l.svcCtx.PushRpc.Subscribe(l.ctx, &rpcReq)
	if err != nil {
		println(err.Error())
	}

	println(rpcRsp.Code, rpcRsp.Msg)

	return &types.TodoResponse{
		Id: id,
	}, nil
}

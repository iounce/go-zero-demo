package logic

import (
	"context"

	"abc.com/todo/v1/internal/svc"
	"abc.com/todo/v1/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var id int = 0

func NewAddTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddTodoLogic {
	return AddTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTodoLogic) AddTodo(req types.TodoRequest) (*types.TodoResponse, error) {
	// todo: add your logic here and delete this line
	id++
	return &types.TodoResponse{
		Id: id,
	}, nil
}

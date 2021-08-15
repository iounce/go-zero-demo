package logic

import (
	"context"
	"strconv"

	"abc.com/todo/v1/internal/svc"
	"abc.com/todo/v1/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetTodoLogic {
	return GetTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTodoLogic) GetTodo(req types.QueryTodoRequest) (*types.QueryTodoResponse, error) {
	// todo: add your logic here and delete this line

	var rsp types.QueryTodoResponse
	size := 2
	rsp.Data = make([]types.QueryTodoOneResponse, size, size)

	for i := 0; i < size; i++ {
		rsp.Data[i].Title = "Title " + strconv.Itoa(i+1)
		rsp.Data[i].Content = "Content " + strconv.Itoa(i+1)
	}

	return &rsp, nil
}

package logic

import (
	"context"

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
	result, err := l.svcCtx.TodoModel.Query(req.Pos, req.Limit)
	if err != nil {
		return nil, err
	}

	size := len(*result)

	var rsp types.QueryTodoResponse

	rsp.Data = make([]types.QueryTodoOneResponse, size, size)

	for i := 0; i < size; i++ {
		rsp.Data[i].Title = (*result)[i].Title
		rsp.Data[i].Content = (*result)[i].Content
	}

	return &rsp, nil
}

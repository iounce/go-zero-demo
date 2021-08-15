package handler

import (
	"net/http"

	"abc.com/todo/v1/internal/logic"
	"abc.com/todo/v1/internal/svc"
	"abc.com/todo/v1/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func addTodoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TodoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddTodoLogic(r.Context(), ctx)
		resp, err := l.AddTodo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

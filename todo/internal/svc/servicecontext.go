package svc

import (
	"rpc/pushclient"

	"abc.com/todo/v1/internal/config"
	"abc.com/todo/v1/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	TodoModel model.TodoModel
	PushRpc   pushclient.Push
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		TodoModel: model.NewTodoModel(conn, c.CacheRedis),
		PushRpc:   pushclient.NewPush((zrpc.MustNewClient(c.PushRpc))),
	}
}

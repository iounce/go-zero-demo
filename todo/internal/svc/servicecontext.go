package svc

import (
	"abc.com/todo/v1/internal/config"
	"abc.com/todo/v1/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	TodoModel model.TodoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		TodoModel: model.NewTodoModel(conn, c.CacheRedis),
	}
}

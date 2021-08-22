package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	todoFieldNames          = builderx.RawFieldNames(&Todo{})
	todoRows                = strings.Join(todoFieldNames, ",")
	todoRowsExpectAutoSet   = strings.Join(stringx.Remove(todoFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	todoRowsWithPlaceHolder = strings.Join(stringx.Remove(todoFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheTodoIdPrefix = "cache:todo:id:"
)

type (
	TodoModel interface {
		Insert(data Todo) (sql.Result, error)
		FindOne(id int64) (*Todo, error)
		Update(data Todo) error
		Delete(id int64) error
		Query(pos int, limit int) (*[]Todo, error)
	}

	defaultTodoModel struct {
		sqlc.CachedConn
		table string
	}

	Todo struct {
		Id      int64  `db:"id"`
		Title   string `db:"title"`
		Content string `db:"content"`
	}
)

func NewTodoModel(conn sqlx.SqlConn, c cache.CacheConf) TodoModel {
	return &defaultTodoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`todo`",
	}
}

func (m *defaultTodoModel) Insert(data Todo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, todoRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Title, data.Content)

	return ret, err
}

func (m *defaultTodoModel) FindOne(id int64) (*Todo, error) {
	todoIdKey := fmt.Sprintf("%s%v", cacheTodoIdPrefix, id)
	var resp Todo
	err := m.QueryRow(&resp, todoIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", todoRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTodoModel) Update(data Todo) error {
	todoIdKey := fmt.Sprintf("%s%v", cacheTodoIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, todoRowsWithPlaceHolder)
		return conn.Exec(query, data.Title, data.Content, data.Id)
	}, todoIdKey)
	return err
}

func (m *defaultTodoModel) Delete(id int64) error {

	todoIdKey := fmt.Sprintf("%s%v", cacheTodoIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, todoIdKey)
	return err
}

func (m *defaultTodoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTodoIdPrefix, primary)
}

func (m *defaultTodoModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", todoRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultTodoModel) Query(pos int, limit int) (*[]Todo, error) {
	var result []Todo
	query := fmt.Sprintf("select %s from %s where (`id` > %d) limit %d",
		todoRows, m.table, pos, limit)
	err := m.QueryRowsNoCache(&result, query)
	switch err {
	case nil:
		return &result, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

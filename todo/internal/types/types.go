// Code generated by goctl. DO NOT EDIT.
package types

type TodoRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TodoResponse struct {
	Id int `json:"id"`
}

type QueryTodoRequest struct {
	Id    string `form:"id,optional"`
	Pos   int    `form:"pos,optional"`
	Limit int    `form:"limit,optional"`
}

type QueryTodoOneResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type QueryTodoResponse struct {
	Data []QueryTodoOneResponse `json:"data"`
}

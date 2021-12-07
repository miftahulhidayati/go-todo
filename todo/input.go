package todo

import "github.com/miftahulhidayati/go-todo/user"

type CreateTodoInput struct {
	Name     string `json:"name" binding:"required"`
	Complete bool   `json:"complete"`
	User     user.User
}
type GetTodoInput struct {
	ID int `uri:"id" binding:"required"`
}

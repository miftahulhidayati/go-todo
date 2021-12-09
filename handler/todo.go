package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miftahulhidayati/go-todo/helper"
	"github.com/miftahulhidayati/go-todo/todo"
	"github.com/miftahulhidayati/go-todo/user"
)

type todoHandler struct {
	service todo.Service
}

func NewTodoHandler(service todo.Service) *todoHandler {
	return &todoHandler{service}
}

// @Tags Todo
// @Summary Get all todo
// @Description Get all todo description
// @Produce json
// @Success 200 {array} helper.Response{data=[]todo.TodoFormatter} "desc"
// @Router /todos [get]
// @Security BearerAuth
func (h *todoHandler) GetTodos(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	todos, err := h.service.GetTodosByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Error to get todos", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of todos", http.StatusOK, "success", todo.FormatTodos(todos))
	c.JSON(http.StatusOK, response)
}

// @Tags Todo
// @Summary Get todo
// @Description Get todo description
// @Produce json
// @Success 200 {array} helper.Response{data=todo.TodoFormatter} "desc"
// @Param id path int true "Todo ID"
// @Router /todos/{id} [get]
// @Security BearerAuth
func (h *todoHandler) GetTodo(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	var input todo.GetTodoInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	todoDetail, err := h.service.GetTodoByID(userID, input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Todo", http.StatusOK, "success", todo.FormatTodo(todoDetail))
	c.JSON(http.StatusOK, response)
}

// @Tags Todo
// @Summary Create todo
// @Description Create todo description
// @Accept json
// @Produce json
// @Param todoId body todo.CreateTodoInputApi true "Create Todo"
// @Success 200 {array} helper.Response{data=todo.TodoFormatter} "desc"
// @Router /todos [post]
// @Security BearerAuth
func (h *todoHandler) CreateTodo(c *gin.Context) {
	var input todo.CreateTodoInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create todo", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTodo, err := h.service.CreateTodo(input)
	if err != nil {
		response := helper.APIResponse("Failed to create todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to create todo", http.StatusCreated, "success", todo.FormatTodo(newTodo))
	c.JSON(http.StatusCreated, response)
}

// @Tags Todo
// @Summary Update todo
// @Description Update todo description
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body todo.CreateTodoInputApi true "Update Todo"
// @Success 200 {array} helper.Response{data=todo.TodoFormatter} "desc"
// @Router /todos/{id} [put]
// @Security BearerAuth
func (h *todoHandler) UpdateTodo(c *gin.Context) {
	var inputID todo.GetTodoInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to get todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData todo.CreateTodoInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update todo", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedTodo, err := h.service.UpdateTodo(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to update todo", http.StatusOK, "success", todo.FormatTodo(updatedTodo))
	c.JSON(http.StatusOK, response)
}

// @Tags Todo
// @Summary delete todo
// @Description delete todo description
// @Produce json
// @Param todoId path int true "Todo ID"
// @Success 200 {array} nil
// @Router /todos/{todoId} [delete]
// @Security BearerAuth
func (h *todoHandler) DeleteTodo(c *gin.Context) {
	var inputID todo.GetTodoInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to get todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	err = h.service.DeleteTodo(userID, inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to delete todo", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

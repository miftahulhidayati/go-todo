package handler

import (
	"net/http"
	"strconv"

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
func (h *todoHandler) GetTodos(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	todos, err := h.service.GetTodosByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Error to get todons", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of todos", http.StatusOK, "success", todo.FormatTodos(todos))
	c.JSON(http.StatusOK, response)
}
func (h *todoHandler) GetTodosByUserID(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	todos, err := h.service.GetTodosByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Error to get todons", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of todos", http.StatusOK, "success", todo.FormatTodos(todos))
	c.JSON(http.StatusOK, response)
}

func (h *todoHandler) GetTodo(c *gin.Context) {
	var input todo.GetTodoInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	todoDetail, err := h.service.GetTodoByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Todo", http.StatusOK, "success", todo.FormatTodo(todoDetail))
	c.JSON(http.StatusOK, response)
}
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
func (h *todoHandler) DeleteTodo(c *gin.Context) {
	var inputID todo.GetTodoInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to get todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.DeleteTodo(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to delete todo", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

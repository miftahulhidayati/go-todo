package handler

import (
	"net/http"

	"github.com/miftahulhidayati/go-todo/auth"
	"github.com/miftahulhidayati/go-todo/helper"
	"github.com/miftahulhidayati/go-todo/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// @Tags User
// @Summary Register user
// @Description Resigter User
// @Accept json
// @Produce json
// @Param user body user.RegisterUserInput true "Register User"
// @Success 200 {array} helper.Response{data=user.UserFormatter} "desc"
// @Router /register [post]
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

// @Tags User
// @Summary Login user
// @Description Login User
// @Accept json
// @Produce json
// @Param userId body user.LoginInput true "Login User"
// @Success 200 {array} helper.Response{data=user.UserFormatter} "desc"
// @Router /login [post]
func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Login account failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(loggedinUser, token)
	response := helper.APIResponse("Successfully loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

// @Tags User
// @Summary Get user
// @Description Get user description
// @Produce json
// @Success 200 {array} helper.Response{data=user.UserFormatter} "desc"
// @Router /users [get]
// @Security BearerAuth
func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	formatter := user.FormatUser(currentUser, "")

	response := helper.APIResponse("Successfully fetch data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

// @Tags User
// @Summary Update user
// @Description Update user description
// @Produce json
// @Param user body user.UpdateUserInput true "Update User"
// @Success 200 {array} helper.Response{data=user.UserFormatter} "desc"
// @Router /users [put]
// @Security BearerAuth
func (h *userHandler) UpdateUser(c *gin.Context) {
	var inputData user.UpdateUserInput

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update user", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedUser, err := h.userService.UpdateUser(userID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(updatedUser.ID)
	if err != nil {
		response := helper.APIResponse("Update account failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to update user", http.StatusOK, "success", user.FormatUser(updatedUser, token))
	c.JSON(http.StatusOK, response)
}

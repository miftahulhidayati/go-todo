package main

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/miftahulhidayati/go-todo/auth"
	"github.com/miftahulhidayati/go-todo/database"
	"github.com/miftahulhidayati/go-todo/handler"
	"github.com/miftahulhidayati/go-todo/helper"
	"github.com/miftahulhidayati/go-todo/todo"
	"github.com/miftahulhidayati/go-todo/user"
)

func main() {
	db := database.InitMysqlDB()

	userRepository := user.NewRepository(db)
	todoRepository := todo.NewRepository(db)

	authService := auth.NewService()
	userService := user.NewService(userRepository)
	todoService := todo.NewService(todoRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	todoHandler := handler.NewTodoHandler(todoService)

	router := gin.Default()
	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))
	router.Use(sessions.Sessions("mysession", cookieStore))

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	api.GET("/todos", todoHandler.GetTodos)
	api.GET("/todos/:id", todoHandler.GetTodo)
	api.POST("/todos", authMiddleware(authService, userService), todoHandler.CreateTodo)
	api.PUT("/todos/:id", authMiddleware(authService, userService), todoHandler.UpdateTodo)
	api.DELETE("/todos/:id", authMiddleware(authService, userService), todoHandler.DeleteTodo)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			respons := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			respons := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			respons := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			respons := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}
		c.Set("currentUser", user)
	}
}

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
	docs "github.com/miftahulhidayati/go-todo/docs"
	"github.com/miftahulhidayati/go-todo/handler"
	"github.com/miftahulhidayati/go-todo/helper"
	"github.com/miftahulhidayati/go-todo/todo"
	"github.com/miftahulhidayati/go-todo/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Todo Application
// @version 1.0
// @description This is a Todo API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email miftahulhdyt@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

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

	// register user
	api.POST("/register", userHandler.RegisterUser)
	// login user
	api.POST("/login", userHandler.Login)
	// get user
	api.GET("/users", authMiddleware(authService, userService), userHandler.FetchUser)
	api.PUT("/users", authMiddleware(authService, userService), userHandler.UpdateUser)

	// get all todos
	api.GET("/todos", authMiddleware(authService, userService), todoHandler.GetTodos)
	// get todo by id
	api.GET("/todos/:id", authMiddleware(authService, userService), todoHandler.GetTodo)
	// create todo
	api.POST("/todos", authMiddleware(authService, userService), todoHandler.CreateTodo)
	// update todo
	api.PUT("/todos/:id", authMiddleware(authService, userService), todoHandler.UpdateTodo)
	// delete todo
	api.DELETE("/todos/:id", authMiddleware(authService, userService), todoHandler.DeleteTodo)

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
			respons := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "accessToken invalid or expired", nil)
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

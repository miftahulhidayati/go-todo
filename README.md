# go-todo 
A RESTful API for todo application with Go

# API Documentation 
https://go-todo-miftahulhdyt.herokuapp.com/swagger/index.html

## Installation & Run
```bash
# Download this project
go get github.com/miftahulhidayati/go-todo
```

Before running API server, you should set the postgre database config, and JWT Secrets with yours on .env
```go
DB_NAME = 
DB_PASS = 
DB_HOST = 
DB_PORT = 
DB_USER = 
JWT_SECRET = 
```
```bash
# Run
cd go-todo
go run main.go
./go-todo
```

# Dependencies
- JWT GO "github.com/dgrijalva/jwt-go"
- GIN "github.com/gin-gonic/gin"
- GODOTENV "github.com/joho/godotenv"
- swaggerFiles "github.com/swaggo/files"
- ginSwagger "github.com/swaggo/gin-swagger"
- GIN CORS "github.com/gin-contrib/cors"
- GIN SESSIONS "github.com/gin-contrib/sessions" "github.com/gin-contrib/sessions/cookie"
- GORM "gorm.io/gorm"

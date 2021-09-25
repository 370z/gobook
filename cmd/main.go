package main

import (
	// "encoding/json"
	"gobook/internal/auth"
	"gobook/internal/post"
	"gobook/internal/user"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// "github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Initialize() {

}

func main() {
	err :=godotenv.Load()
	if err!=nil {
		log.Fatal("Error loading .env")
	}
	e := echo.New()
	// Connect to db
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&post.Post{})

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authService := auth.NewService(db)
	userService := user.NewService(db)
	postService := post.NewService(db)

	//Router
	e.GET("/public", func(c echo.Context) error {
		return c.String(http.StatusOK, "accessible")
	})
	//Private
	r := e.Group("/private")
	config := middleware.JWTConfig{
		Claims:     &auth.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "private naja")
	})
	// Post Service
	r.POST("/posts", postService.Create)
	//CRUD

	e.POST("/login", authService.Login)
	e.POST("/users", userService.CreateUser)
	e.GET("/users", userService.GetUsers)
	e.GET("/user/:id", userService.GetUser)
	e.PUT("/users/:id", userService.UpdateUser)
	e.DELETE("/users/:id", userService.DeleteUser)

	//localhost:1323
	e.Logger.Fatal(e.Start(":1323"))
}

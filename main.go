package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"todo_app_api/todo"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	db.AutoMigrate(&todo.Todo{})

	todoRepository := todo.CreateNewRepository(db)
	todoService := todo.CreateNewService(todoRepository)
	todoHandler := todo.CreateNewHandler(todoService)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	r.GET("/todos", todoHandler.GetAll)
	r.GET("/todos/:id", todoHandler.GetById)
	r.POST("/todos", todoHandler.AddTodo)
	r.PUT("/todos/:id", todoHandler.UpdateTodo)
	r.PATCH("/todos/:id", todoHandler.ToggleIsCompleteTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)

	r.Run()
}

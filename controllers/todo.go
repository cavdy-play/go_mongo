package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Completed string    `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllTodos(c *gin.Context) {
}

func CreateTodo(c *gin.Context) {
}

func GetSingleTodo(c *gin.Context) {
}

func EditTodo(c *gin.Context) {
}

func DeleteTodo(c *gin.Context) {
}

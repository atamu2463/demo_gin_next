package handlers

import (
	"net/http"

	"github.com/atamu2463/demo_gin_next/domain"
	"github.com/atamu2463/demo_gin_next/repositories"
	"github.com/gin-gonic/gin"
)

// Todoを作成
func CreateTodo(c *gin.Context) {
	var todo domain.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repositories.DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

// Todo一覧を取得
func GetTodos(c *gin.Context) {
	var todos []domain.Todo

	repositories.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

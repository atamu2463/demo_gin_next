package handlers

import (
	"net/http"

	"github.com/atamu2463/demo_gin_next/domain"
	"github.com/atamu2463/demo_gin_next/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Todoを作成
func CreateTodo(c *gin.Context) {
	var todo domain.Todo
	var db *gorm.DB
	repositories.InitDB()

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

// Todo一覧を取得
func GetTodos(c *gin.Context) {
	var todos []domain.Todo
	var db *gorm.DB
	repositories.InitDB()

	db.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

package main

import (
	"github.com/atamu2463/demo_gin_next/handlers"
	"github.com/atamu2463/demo_gin_next/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	//Ginのコアを生成
	r := gin.Default()

	repositories.InitDB()

	//ルーティング設定
	r.POST("/todos", handlers.CreateTodo)
	r.GET("/todos", handlers.GetTodos)

	//サーバーを起動
	r.Run(":8080")
}

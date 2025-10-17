package main

import (
	"github.com/atamu2463/demo_gin_next/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	//Ginのコアを生成
	r := gin.Default()

	repositories.InitDB()

	// //goodと送るとmorningと返すAPIを作成
	// r.GET("/good", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "morning",
	// 	})
	// })

	//サーバーを起動
	r.Run(":8080")
}

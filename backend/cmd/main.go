package main

import "github.com/gin-gonic/gin"

func main() {
	//Ginのコアを生成
	r := gin.Default()

	//goodと送るとmorningと返すAPIを作成
	r.GET("/good", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "morning",
		})
	})

	//サーバーを起動
	r.Run(":8080")
}

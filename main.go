package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func routes() *gin.Engine {
	r := gin.Default()
	r.GET("/", showRoot)
	return r
}

func showRoot(c *gin.Context){
	message := "message"
	data := "pong"
	c.JSON(200, gin.H{
		message: data,
	})
}

func main() {
	r := routes()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
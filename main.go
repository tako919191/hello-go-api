package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.Run(":3000")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":  "pong",
		"message2": "pang",
	})
}

package main

import (
	"chat-service/infrastructure"

	"github.com/gin-gonic/gin"
)



func main() {
	// ctx := context.Background()
	go infrastructure.Hub.Run()

	router := gin.New()

	router.GET("/room/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		infrastructure.ServeWs(c.Writer, c.Request, roomId)
	})

	router.Run(":80")
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
		r := gin.Default()
		handler := Handler{}

		r.POST("/parse", handler.HandleParseVideo)
		r.GET("/video", handler.HandleGetVideo)
		r.Run()
}

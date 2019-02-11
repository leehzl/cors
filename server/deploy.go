package main

import (
	"github.com/gin-gonic/gin"

	"lilynlee.com/cors/handler"
	"lilynlee.com/cors/middlewares"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(middlewares.FilterMiddleware)
	router.GET("/simple", handler.Simple)
	router.POST("/options", handler.Options)
	router.GET("/cookie", handler.Cookies)
	router.Run(":8080")
}

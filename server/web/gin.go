package web

import (
	"github.com/gin-gonic/gin"
	"lilynlee.com/cors/handler"
	"lilynlee.com/cors/middlewares"
)

func RunGin(addr string) {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	initRouter(engine)

	go engine.Run(addr)
}

func initRouter(engine *gin.Engine) {
	v1 := engine.Group("/v1", middlewares.FilterMiddleware)
	{
		v1.GET("/cors", handler.Cors)
	}
}

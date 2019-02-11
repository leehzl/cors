package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FilterMiddleware(c *gin.Context) {
	fmt.Println("before middleware")
	c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	//c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == "OPTIONS" {
		fmt.Println("options")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.JSON(http.StatusOK, "SUCCESS")

		return
	}
	c.Next()
	fmt.Println("after middleware")
}

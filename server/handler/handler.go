package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Simple(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func Options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func Cookies(c *gin.Context) {
	school, err := c.Cookie("school")
	if err != nil {
		c.SetCookie("school", "cug", 30, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
	if school == "" {
		c.SetCookie("school", "cug", 30, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"school":  school,
	})
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 设置路由组，方便管理
	userGroup := r.Group("/user")
	{ // {}非必须，写上逻辑更清晰
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"path": "user/path"})
		})
		userGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"path": "user/login"})
		})
		userGroup.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"path": "user/postlogin"})
		})

	}

	r.Run()
}

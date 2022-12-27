package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func index0(c *gin.Context) {
	c.String(200, "这是index0()\n")
}

func index1(c *gin.Context) {
	c.String(200, "这是index1()\n")
	c.Abort()
}
func index2(c *gin.Context) {
	c.String(200, "这是index2()\n")
}

func m1(c *gin.Context) {
	fmt.Println("m1 in")
	start := time.Now()
	c.Next() // 到这就去执行其后面的函数，执行完在接着执行下面的语句
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out……")
}
func _07main() {
	router := gin.Default()
	defer router.Run(":8080")
	router.Use(m1) // 注册全局的中间件后就不用在每个函数里都写了（中间件个数不限）
	//router.GET("/0", m1, index2)
	//router.GET("/1", m1, index2)
	//router.GET("/2", m1, index2)
	router.GET("/0", index0)
	router.GET("/1", index1)
	router.GET("/2", index2)

	// 访问不存在的路由时
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"提示：": "页面不存在"})
	})

}

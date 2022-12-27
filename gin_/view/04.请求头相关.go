package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func _04main() {
	router := gin.Default()
	//启动监听，gin将web服务运行在本机的8080端口
	defer router.Run("0.0.0.0:8080")
	router.GET("/", func(c *gin.Context) {
		// 首字母大小写不区分  单词与单词之间用 - 连接
		// 用于获取一个请求头
		fmt.Println(c.GetHeader("User-Agent"))
		fmt.Println(c.GetHeader("user-agent"))
		fmt.Println(c.GetHeader("user-Agent"))
		fmt.Println(c.GetHeader("user-AGent"))

		// Header 是一个普通的 map[string][]string
		//fmt.Println(c.Request.Header)
		// 如果是使用 Get方法或者是 .GetHeader,那么可以不用区分大小写，并且返回第一个value
		fmt.Println(c.Request.Header.Get("User-Agent"))

		// 如果是用map的取值方式，请注意大小写问题
		fmt.Println(c.Request.Header["User-Agent"])
		fmt.Printf("类型=%T\n", c.Request.Header)
		fmt.Println(c.Request.Header["user-agent"])

		// 自定义的请求头，用Get方法也是免大小写
		fmt.Println(c.Request.Header.Get("Token"))
		//fmt.Println(c.Request.Header.Get("token"))
		c.JSON(200, gin.H{"msg": "成功"}) //  返回给浏览器
	})

	//识别爬虫的简单原理
	router.GET("/index", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		if strings.Contains(userAgent, "python") {
			c.JSON(0, gin.H{"data": "你是爬虫"})
			return
		}
		c.JSON(0, gin.H{"data": "你是正常用户"})
	})

	// 设置响应头
	router.GET("/res", func(c *gin.Context) {
		// 可以自定义响应头
		c.Header("Content-Type", "application/xml; charset=utf-8")
		c.Header("zz", "自定义响应头")
		c.JSON(0, gin.H{"data": "看响应头！"})
	})
}

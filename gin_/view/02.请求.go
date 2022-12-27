package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func _query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.QueryArray("user")) // 拿到多个相同的查询参数(返回值是数组)
	fmt.Println(c.DefaultQuery("addr", "四川省"))
}

func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

func _form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("addr", "山东省"))
	forms, err := c.MultipartForm() // 接收所有的form表单参数，包括文件
	fmt.Println(forms, err)
}

// 原始参数
func _raw(c *gin.Context) {
	body, _ := c.GetRawData()
	fmt.Println(body)
	fmt.Println(string(body))
}
func __main() {
	router := gin.Default()
	defer router.Run(":8080")
	// 查询参数query
	router.GET("/query", _query) // 192.168.220.1:8080/query?id=1&user=Bob

	// 动态参数
	router.GET("/param/:user_id/", _param)         // 192.168.220.1:8080/param/123  获取到param后面的123
	router.GET("/param/:user_id/:book_id", _param) // // 192.168.220.1:8080/param/123/456  获取到param后面的456

	// 表单参数
	router.POST("/form", _form)

	// 原始参数
	router.POST("/raw", _raw)

}

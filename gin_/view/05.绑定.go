package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name string `json:"name"form:"name" uri:"name" binding:"startswith=like"`
	Age  int    `json:"age" form:"age"  uri:"age" `
	Sex  string `json:"sex" form:"sex"  uri:"sex" binding:"oneof=男 女" `
	//LikeList []string `json:"likeList" binding:"dive, startswith=like" `
}

func _05main() {
	router := gin.Default()
	defer router.Run(":8080")
	// json格式绑定，http://127.0.0.1:8080/，是在body里直接传json格式数据
	router.POST("/", func(c *gin.Context) {

		var userInfo UserInfo
		err := c.ShouldBindJSON(&userInfo)
		if err != nil {
			//fmt.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, userInfo)

	})

	// 查询绑定，http://127.0.0.1:8080/query?name=ss&age=2&sex=男
	router.POST("/query", func(c *gin.Context) {

		var userInfo UserInfo
		err := c.ShouldBindQuery(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, userInfo)
	})

	// 绑定动态参数  http://127.0.0.1:8080/uri/宋旭辉/12/男（是post请求哦，不能直接在浏览器输入）
	router.POST("/uri/:name/:age/:sex", func(c *gin.Context) {

		var userInfo UserInfo
		err := c.ShouldBindUri(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, userInfo)

	})

	// 对form-data和xxx-www-form-urlencoded数据进行绑定
	router.POST("/form", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBind(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, userInfo)
	})
}

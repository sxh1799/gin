package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func _main() {
	fmt.Println("hello world!")
	// 创建默认路由
	router := gin.Default()

	//启动监听，gin将web服务运行在本机的8080端口
	defer router.Run("0.0.0.0:8080")

	// 绑定路由规则和对应的处理函数（即访问index时，就由相应的函数去处理）
	router.GET("/index", func(context *gin.Context) {
		context.String(200, "Hello world你好")
	})

	router.GET("/json", func(context *gin.Context) {
		//直接返回json数据，无需转换
		context.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	// 结构体转json
	router.GET("/moreJSON", func(context *gin.Context) {
		// You also can use a struct
		type Msg struct {
			Name    string `json:"user"` // 转换为json时会Name会变为user
			Message string
			Number  int `json:"-"` // `json:"-"`表示不把这个字段转为json，比如密码之类的隐私数据
		}
		msg := Msg{"Jack", "hey", 21}
		// 注意 msg.Name 变成了 "user" 字段
		// 以下方式都会输出 :   {"user": "hanru", "Message": "hey", "Number": 123}
		context.JSON(http.StatusOK, msg)
	})
	// map转json
	router.GET("/map", func(context *gin.Context) {
		userMap := map[string]string{
			"userName": "Bob",
			"age":      "32",
		}
		context.JSON(200, userMap)
	})

	// 相应xml
	router.GET("/xml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"user": "hanru", "message": "hey", "status": http.StatusOK})
	})

	// 响应yaml
	router.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"user": "hanru", "message": "hey", "status": http.StatusOK})
	})

	// 响应html
	//加载模板
	router.LoadHTMLGlob("template/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	//定义路由
	router.GET("/tem", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "02.baidu.html", gin.H{})
	})

	// 响应文件
	router.StaticFile("/qqzone", "static/QQzone.png")    // 只允许看到这个目录下的指定文件
	router.StaticFS("static", http.Dir("static/static")) // 可以看到指定目录下的所有文件

	// 重定向（即跳转网页）
	router.GET("/redirect", func(c *gin.Context) {
		//支持内部和外部的重定向 301是永久重定向，302是临时的，也可以定向到本地网址
		c.Redirect(302, "https://www.google.com/")
	})

}

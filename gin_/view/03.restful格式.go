package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

import (
	"encoding/json"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"` // any泛型，空接口
	Msg  string `json:"msg"`
}

func _bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	fmt.Println("body=", string(body))
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

// _getList 文章列表页面
func _getList(c *gin.Context) {
	// 包含搜索，分页
	articleList := []ArticleModel{
		{"Go语言入门", "这篇文章是《Go语言入门》"},
		{"python语言入门", "这篇文章是《python语言入门》"},
		{"JavaScript语言入门", "这篇文章是《JavaScript语言入门》"},
	}
	c.JSON(200, Response{0, articleList, "成功"})
}

// _getDetail 文章详情
func _getDetail(c *gin.Context) {
	// 获取param中的id
	fmt.Println(c.Param("id"))
	article := ArticleModel{
		"Go语言入门",
		"这篇文章是《Go语言入门》",
	}
	c.JSON(200, Response{0, article, "成功"})
}

// _create 创建文章
func _create(c *gin.Context) {
	// 接收前端传递来的json数据
	var article ArticleModel

	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, Response{0, article, "添加成功"})
}

// _update 编辑文章
func _update(c *gin.Context) {
	fmt.Println(c.Param("id"))
	var article ArticleModel
	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, Response{0, article, "修改成功"})
}

// _delete 删除文章
func _delete(c *gin.Context) {
	fmt.Println(c.Param("id"))
	c.JSON(200, Response{0, map[string]string{}, "删除成功"})
}

func ___main() {
	router := gin.Default()
	router.GET("/articles", _getList)       // 文章列表
	router.GET("/articles/:id", _getDetail) // 文章详情
	router.POST("/articles", _create)       // 添加文章
	router.PUT("/articles/:id", _update)    // 编辑文章
	router.DELETE("/articles/:id", _delete) // 删除文章
	router.Run(":8080")
}

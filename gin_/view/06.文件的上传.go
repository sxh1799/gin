package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func _06main() {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// 单位是字节， << 是左移预算符号，等价于 8 * 2^20
	// gin对文件上传大小的默认值是32MB
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")

		fileRead, _ := file.Open()
		// 读取文件内容
		data, err := io.ReadAll(fileRead)
		fmt.Println("文件内容：" + string(data))
		dst := "./" + file.Filename

		// 创建一个文件
		out, err := os.Create(dst)
		if err != nil {
			fmt.Println(err)
		}
		defer out.Close()
		// 拷贝文件对象到out中
		io.Copy(out, fileRead)

		fmt.Println("文件名：" + file.Filename)
		fmt.Printf("文件大小：%d%s\n", file.Size>>10, "kb")

		// 设置文件保存的路径
		dst = "./upload/" + file.Filename
		// 保存文件
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			fmt.Println("err=", err)
		}
		// 在浏览器端输出
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// 同时上传多个文件
	router.POST("/uploads", func(c *gin.Context) {
		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println("err=", err)
		}
		files := form.File["file[]"] // 注意这里名字不要对不上了

		for _, file := range files {
			log.Println(file.Filename)
			// 上传文件至指定目录
			c.SaveUploadedFile(file, "./"+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	// 下载文件
	router.GET("/download", func(c *gin.Context) {
		c.Header("Content-Type", "application/octet-stream")               // 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
		c.Header("Content-Disposition", "attachment; filename="+"牛逼1.txt") // 用来指定下载下来的文件名

		c.File("./1.txt")
	})
	router.Run(":8080")
}

// 6-1-upload-file.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1创建路由,默认使用了两个中间件Logger(),Recovery()
	r := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	// r.MaxMultipartMemory = 8 << 20  // 8 MiB
	// 指明html加载文件目录
	r.LoadHTMLGlob("./*")
	r.Handle("GET", "/", func(context *gin.Context) {
		fmt.Println("/访问")
		// 返回HTML文件，响应状态码200，html文件名为index.html，模板参数为nil
		context.HTML(http.StatusOK, "6-1-upload-file.html", nil)
	})
	// 2绑定路由规则,
	// gin.Context,封装了request和respose
	r.POST("/upload", func(c *gin.Context) {
		

		file, _ := c.FormFile("file")
		log.Println("file:", file.Filename)
		c.SaveUploadedFile(file, file.Filename) // 上传文件到指定的路径
		c.String(200, fmt.Sprintf("%s upload file!", file.Filename))
	})
	// 3监听端口，默认8080
	r.Run(":8080")
}

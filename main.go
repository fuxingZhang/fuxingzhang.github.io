package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//router.StaticFS("/", http.Dir(filepath.Join(dir, "./html")))

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, TLS user! Your config: %s", c.Request.TLS)
	})

	r.GET("/", func(c *gin.Context) {
		c.File("./html/index.html")
	})

	r.Static("/static", "./static")

	// 遍历html目录的一级深度
	files, err := os.ReadDir("./html")
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".html") {
			// 构建URL路径，去除".html"扩展名
			urlPath := "/" + strings.TrimSuffix(file.Name(), ".html")

			// 注册路由
			r.GET(urlPath, func(c *gin.Context) {
				// 提供HTML文件
				c.File(filepath.Join("./html", file.Name()))
			})
		}
	}

	go func() {
		r.Run(":80")
	}()

	log.Fatal(autotls.Run(r, "zhangfuxing.icu"))
}

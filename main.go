package main

import (
	"app/pkg/controller"
	"app/pkg/router"
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())

	//r.StaticFS("/", http.Dir(filepath.Join(dir, "./html")))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/cookie", func(c *gin.Context) {
		controller.HandleProxy(c)
		c.Abort()
	})

	r.GET("/", func(c *gin.Context) {
		c.File("./html/index.html")
	})

	r.Static("/static", "./static")

	err := router.RegisterHTMLRoutes(r, "./html", "/")
	if err != nil {
		log.Fatalf("Failed to register HTML routes: %v", err)
	}

	go func() {
		r.Run(":80")
	}()

	log.Fatal(autotls.Run(r, "hainan888.top", "stockapp.sandhope.com"))
}

package main

import (
	"app/headless/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/cookie", func(c *gin.Context) {
		var domain = c.Query("domain")
		var key = c.Query("key")

		if domain == "" {
			c.AbortWithStatus(400)
			return
		}

		result, err := pkg.GetCookie(domain)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		if key != "" {
			c.String(http.StatusOK, result[key])
			return
		}

		c.JSON(http.StatusOK, result)
	})

	r.Run(":8090")
}

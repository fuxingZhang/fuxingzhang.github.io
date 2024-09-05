package controller

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

var target, _ = url.Parse("http://127.0.0.1:8090/cookie")

func HandleProxy(c *gin.Context) {
	proxy := httputil.ReverseProxy{
		ErrorLog: log.Default(),
		Director: newDirector(target),
	}
	proxy.ServeHTTP(c.Writer, c.Request)
	c.Abort()
}

func newDirector(target *url.URL) func(*http.Request) {
	targetQuery := target.RawQuery
	return func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
		req.Host = target.Host
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}
}

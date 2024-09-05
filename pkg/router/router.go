package router

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterHTMLRoutes(router *gin.Engine, baseDir string, prefix string) error {
	files, err := os.ReadDir(baseDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		path := filepath.Join(baseDir, file.Name())
		if file.IsDir() {
			err := RegisterHTMLRoutes(router, path, prefix+file.Name()+"/")
			if err != nil {
				return err
			}
		} else if strings.HasSuffix(strings.ToLower(file.Name()), ".html") {
			urlPath := prefix + strings.TrimSuffix(file.Name(), ".html")
			router.GET(urlPath, func(c *gin.Context) {
				c.File(path)
			})
		}
	}
	return nil
}

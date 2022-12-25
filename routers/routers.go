package routers

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/seaung/vhub/pkg/api"
)

func renderTemplate(templateDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	base, err := filepath.Glob(templateDir + "/public/base.html")
	if err != nil {
		panic(err.Error())
	}

	vulsDir, err := filepath.Glob(templateDir + "/pages/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _, dir := range vulsDir {
		layoutdir := make([]string, len(base))
		copy(layoutdir, vulsDir)
		files := append(layoutdir, dir)
		r.AddFromFiles(filepath.Base(dir), files...)
	}

	return r
}

func InitRouters() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(gin.Logger(), gin.Recovery())
	r.HTMLRender = renderTemplate("./templates")

	r.LoadHTMLGlob("templates/**/*.html")
	r.Static("/static", "./static")

	r.GET("/index", api.HomePage)
	r.GET("/rce", api.RecController)
	return r
}

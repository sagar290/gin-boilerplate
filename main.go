package main

import (
	router "gin-boilerplate/Router"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	//new template engine
	r.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "views",
		Extension:    ".html",
		Master:       "index",
		Partials:     []string{"partials/ad"},
		DisableCache: true,
	})

	// register home route
	router.RegisterHomeRoute(r)

	r.Run(":8000")
}

package main

import (
	"fmt"

	"gin-boilerplate/Controller"

	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Hello world")

	r := gin.Default()
	//new template engine
	r.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "views",
		Extension:    ".html",
		Master:       "index",
		Partials:     []string{"partials/ad"},
		DisableCache: true,
	})

	r.Static("/assets", "./public/assets")

	r.GET("/", Controller.HomePage)
	r.POST("/PostHomePage", Controller.PostHomePage)
	r.GET("/query", Controller.QueryString)
	r.GET("/param/:name/:age", Controller.ParamString)

	r.Run(":8000")
}

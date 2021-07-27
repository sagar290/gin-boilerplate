package main

import (
	"database/sql"
	"fmt"
	"time"

	"gin-boilerplate/Controller"
	"gin-boilerplate/Model"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	// Result := map[string]interface{}{}

	fmt.Println(Model.Host)
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

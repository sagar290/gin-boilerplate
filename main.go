package main

import (
	"fmt"

	"Controller"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Hello world")

	r := gin.Default()
	r.GET("/HomePage", Controller.HomePage)
	r.POST("/PostHomePage", Controller.PostHomePage)
	r.GET("/query", Controller.QueryString)
	r.GET("/param/:name/:age", Controller.ParamString)

	r.Run(":8000")
}

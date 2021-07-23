package Controller

import "github.com/gin-gonic/gin"

type Formbody struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func PostHomePage(c *gin.Context) {
	var formbody Formbody

	c.BindJSON(&formbody)

	c.JSON(200, gin.H{
		"name": formbody.Name,
		"age":  formbody.Age,
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func ParamString(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

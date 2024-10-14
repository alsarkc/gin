package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Demoinfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Getdemo(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	fmt.Println(name, age)
}

// 方法1(推荐)
func Postdemo(c *gin.Context) {
	var demoinfo Demoinfo
	c.BindJSON(&demoinfo)
	fmt.Println(demoinfo)
}

// 方法2
func Postdemo2(c *gin.Context) {
	var jsonData map[string]interface{}
	c.BindJSON(&jsonData)
	fmt.Println(jsonData)
}

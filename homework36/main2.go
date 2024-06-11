package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", NewFunc)
	route.Run(":8088")
}

func NewFunc(c *gin.Context) {
	var person Person
	err := c.ShouldBindUri(&person);
	if  err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
}
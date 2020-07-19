package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func init() {
	initDB()
}

func main() {
	server := gin.Default()

	server.GET("/", listPeople)
	server.POST("/", createPeople)

	server.Run(":8086")
}

var biz = newUserBiz()

func listPeople(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	cond := make(map[string]interface{})
	if ageCond := c.Query("age"); ageCond != "" {
		cond["age"] = ageCond
	}

	users, err := biz.Query(c, cond)
	if err != nil {
		c.JSON(500, gin.H{"err": err})
	} else {
		c.JSON(200, gin.H{"data": users})
	}
}

func createPeople(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	var u user
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
		c.JSON(500, gin.H{"err": err})
	}

	if err := biz.Craete(c, u); err != nil {
		c.JSON(500, gin.H{"err": err})
		return
	}

	c.JSON(200, nil)
}

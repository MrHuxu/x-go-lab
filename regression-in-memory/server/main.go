package main

import (
	"encoding/json"

	"github.com/MrHuxu/x-go-lab/regression-in-memory/server/models"

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

	users, err := biz.Query(c, c.Query("age"))
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
	} else {
		c.JSON(200, gin.H{"data": users})
	}
}

func createPeople(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	var u models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
		c.JSON(500, gin.H{"err": err})
	}

	if err := biz.Craete(c, &u); err != nil {
		c.JSON(500, gin.H{"err": err})
		return
	}

	c.JSON(200, nil)
}

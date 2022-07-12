package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/salemzii/utcconverter/parser"
)

func main() {
	fmt.Println(time.Now())
	router := gin.Default()

	router.GET("/", welcome)
	router.POST("/convert", RecieveDateTime)

	router.Run()

}

func welcome(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello welcome to utc converter webhook",
	})
}

type DateTimeReq struct {
	Datetime string `json:"datetime"`
}

func RecieveDateTime(c *gin.Context) {
	var tm DateTimeReq

	if c.Request.Method == "POST" {
		if err := c.ShouldBindJSON(&tm); err != nil {
			log.Println("Error ", err)
		}
	}
	loc, _ := time.LoadLocation("Australia/Sydney")
	t := parser.Parse(tm.Datetime)
	c.JSON(200, gin.H{
		"datetime": t.In(loc).String(),
	})
}

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.New()
	router.GET("/someXML",func(c *gin.Context){
		type MessageRecord struct {
			Name     string   `xml:"user"` 
			Message  string 
			Number   int
		}

		var msg MessageRecord
		msg.Name = "Lena"
		msg.Message = "Hello World"
		msg.Number = 1341431243
		c.XML(http.StatusOK, msg)
	})
	router.Run()
}
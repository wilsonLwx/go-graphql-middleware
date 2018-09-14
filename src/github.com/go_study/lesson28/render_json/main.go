package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/someJSON",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{"message": "hello world","status": http.StatusOK})
	})
	router.GET("/moreJSON",func(c *gin.Context){
		var msg struct{
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hello world"
		msg.Number = 12532
		c.JSON(http.StatusOK, msg)
	})
	router.Run()
}
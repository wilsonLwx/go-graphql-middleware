package main

import (
	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()
	r.GET("/user/info/:username/:address/",func(c *gin.Context){
		username := c.Param("username")
		address := c.Param("address")
	// r.GET("/user/info",func(c *gin.Context){
		// username := c.DefaultQuery("username","laowang")
		// address := c.Query("address")

		c.JSON(200,gin.H{
			"message":"pong",
			"username":username,
			"address":address,
		})
	})
	r.Run()
}
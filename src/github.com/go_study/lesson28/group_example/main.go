package main

import (
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context){
	c.JSON(200,gin.H{
		"message": "login",
	})
}
func submit(c *gin.Context){
	c.JSON(200,gin.H{
		"message": "submit",
	})
}
func read(c *gin.Context){
	c.JSON(200,gin.H{
		"message": "read",
	})
}
func main(){
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/login",login)
		v1.POST("/submit",submit)
		v1.POST("/read",read)
	}
	v2 := router.Group("/v2")
	{
		v2.POST("/login",login)
		v2.POST("/submit",submit)
		v2.POST("/read",read)
	}
router.Run(":8080")
}
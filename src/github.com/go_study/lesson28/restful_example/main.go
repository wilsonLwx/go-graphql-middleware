package main


import (
	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()
	r.GET("/user/info",func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"this is get method!",
		})
	})
	r.POST("/user/info",func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"this is post method!",
		})
	})
	r.PUT("/user/info",func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"this is put method!",
		})
	})
	r.DELETE("/user/info",func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"this is delete method!",
		})
	})
	r.Run()
}

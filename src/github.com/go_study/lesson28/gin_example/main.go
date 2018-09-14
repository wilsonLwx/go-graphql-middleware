package mian

import "github.com/gin-gonic/gin"

func handleFunc(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"Hello World!",
	})
}

func main(){
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping",func(c *gin.Context){
		// 输出json结果给调用者
		c.JSON(200,gin.H{
			"message":"pong",
		})
	})
	r.GET("/test",handleFunc)
	r.Run(":9090")//listen and serve on 0.0.0.0:8080
}



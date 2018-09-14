package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


type Login struct{
	User string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main(){
	router := gin.Default()
	router.POST("/loginJSON",func(c *gin.Context){
		var login Login
		if err := c.ShouldBindJSON(&login); err == nil{
			fmt.Printf("login info:%+v\n",login)
			c.JSON(http.StatusOK,gin.H{
				"user": login.User,
				"password": login.Password,
			})
		}else{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		}
	})

	router.POST("/loginForm",func(c *gin.Context){
		var login Login
		if err := c.ShouldBind(&login); err == nil{
			fmt.Printf("login info:%#v\n",login)
			c.JSON(http.StatusOK,gin.H{
				"user": login.User,
				"password": login.Password,
			})
		}else{
			c.JSON(http.StatusBadRequest,gin.H{"errot":err.Error()})
		}
	})

	router.GET("loginForm",func(c *gin.Context){
		var login Login
		if err := c.ShouldBind(&login); err == nil{
			c.JSON(http.StatusOK,gin.H{
				"user": login.User,
				"password": login.Password,
			})
		}else{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		}
	})
	router.Run()
}
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.POST("/upload",func(c *gin.Context){
		file,err := c.FormFile("file")
		if err != nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"message":err.Error(),
			})
			return
		}
		log.Println(file.Filename)
		dst := fmt.Sprintf("./%s",file.Filename)
		c.SaveUploadedFile(file,dst)
		c.JSON(http.StatusOK,gin.H{
			"message":fmt.Sprintf("'%s' upload!",file.Filename),
		})
	})
	router.Run(":8080")
}
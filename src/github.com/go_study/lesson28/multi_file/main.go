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
		form, _ := c.MultipartForm()
		files := form.File["file"]
		for index, file := range files{
			log.Println(file.Filename)
			dst := fmt.Sprintf("./%s_%d",file.Filename,index)
			c.SaveUploadedFile(file,dst)
		}
		c.JSON(http.StatusOK,gin.H{
			"message":fmt.Sprintf("%d file upload!",len(files)),
		})
	})
	router.Run(":8080")
}

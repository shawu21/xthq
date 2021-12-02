package main

import (
	"Test/controller"
	"Test/mySqldb"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	bGroup := r.Group("bolog")
	{
		bGroup.GET("/get/article", controller.GetArticleByID)
		bGroup.GET("/get/comment", controller.GetComment)
		bGroup.POST("/post/atricle", controller.UploadArticle)
		bGroup.POST("/post/comment", controller.UploadComment)
	}
	r.Run()
	defer mySqldb.MySqlDb.Close()
}

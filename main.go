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
		bGroup.GET("/get/articleid", controller.GetArticleByID)
		bGroup.GET("/get/articletag", controller.GetArticleByTag)
		bGroup.GET("/get/articlecategory", controller.GetArticleByCategory)

		bGroup.GET("/get/comments", controller.GetComment) //查询
		bGroup.GET("/get/comment", controller.GetCommentByID)

		bGroup.POST("/post/article", controller.UploadArticle)
		bGroup.POST("/post/comment", controller.UploadComment)
		bGroup.POST("/post/reply", controller.ReplyComment)
	}
	r.Run()
	defer mySqldb.MySqlDb.Close()
}

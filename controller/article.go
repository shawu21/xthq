package controller

import (
	"Test/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadArticle(c *gin.Context) {
	var artcile model.ArticleInfo
	artcile.Content = c.PostForm("content")
	artcile.Tag = c.PostForm("tag")
	artcile.Category = c.PostForm("category")
	model.SaveArtical(artcile)
	c.JSON(http.StatusOK, gin.H{
		"content":  artcile.Content,
		"category": artcile.Category,
		"tag":      artcile.Tag,
	})
} //上传文章

func GetArticleByID(c *gin.Context) {
	ID := c.Query("id")
	Id, _ := strconv.Atoi(ID)
	content := model.FindArtical(Id)
	c.JSON(http.StatusOK, gin.H{
		"articleid": Id,
		"content":   content,
	})
} //通过ID查询文章，返回内容

func GetArticleByCategory(c *gin.Context) {
	cate := c.Query("category")
	artciles := model.FindArticalIDbyCate(cate)
	c.JSON(http.StatusOK, model.ApiReturn(artciles.Msg, artciles.Data))
} //通过分类category查询文章，返回文章ID

func GetArticleByTag(c *gin.Context) {
	tag := c.Query("tag")
	articles := model.FindArticalIDbytag(tag)
	c.JSON(http.StatusOK, model.ApiReturn(articles.Msg, articles.Data))
} //通过tag查询文章，返回文章ID

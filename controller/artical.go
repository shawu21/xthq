package controller

import (
	"Test/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadArticle(c *gin.Context) {
	var artcile model.ArticleInfo
	artcile.Content = c.Query("content")
	artcile.Tag = c.Query("tag")
	artcile.Category = c.Query("category")
	model.SaveArtical(artcile)
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
	c.JSON(http.StatusOK, gin.H{
		"these are artcile's id about": cate,
	})
	for _, a := range artciles {
		fmt.Println(a.ID)
	}
} //通过分类category查询文章，返回文章ID

func GetArticleByTag(c *gin.Context) {
	tag := c.Query("tag")
	articles := model.FindArticalIDbytag(tag)
	c.JSON(http.StatusOK, gin.H{
		"these are artcile's id about": tag,
	})
	for _, a := range articles {
		fmt.Println(a.ID)
	}
} //通过tag查询文章，返回文章ID

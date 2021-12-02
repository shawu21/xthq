package controller

import (
	"Test/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadArticle(c *gin.Context) {

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
	c.JSON(http.StatusOK, gin.H{
		"these are artcile's id about": cate,
	})
	model.FindArticalIDbyCate(cate)
} //通过分类category查询文章，返回文章ID

func GetArticleByTag(c *gin.Context) {
	tag := c.Query("tag")
	c.JSON(http.StatusOK, gin.H{
		"these are artcile's id about": tag,
	})
	model.FindArticalIDbytag(tag)
} //通过tag查询文章，返回文章ID

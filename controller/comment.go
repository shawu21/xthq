package controller

import (
	"Test/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadComment(c *gin.Context) {

} //对文章进行评论，需要指明ID，同时将该评论附上ID,且留下邮箱

func JundgeEmail(c *gin.Context) {

} //判断是否留下邮箱

func ReplyComment(c *gin.Context) {

} //对评论进行回复，需要指明评论的ID，同时发送邮件表明发送成功

func GetComment(c *gin.Context) {
	ID := c.Query("id")
	Id, _ := strconv.Atoi(ID)
	content := model.FindComment(Id)
	c.JSON(http.StatusOK, gin.H{
		"commentid": Id,
		"content":   content,
	})
} //查询评论

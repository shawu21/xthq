package controller

import (
	"github.com/gin-gonic/gin"
)

func Comment(c *gin.Context) {

} //对文章进行评论，需要指明ID，同时将该评论附上ID,且留下邮箱

func JundgeEmail(c *gin.Context) {

} //判断是否留下邮箱

func ReplyComment(c *gin.Context) {

} //对评论进行回复，需要指明评论的ID，同时发送邮件表明发送成功

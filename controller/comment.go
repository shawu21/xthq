package controller

import (
	"Test/model"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadComment(c *gin.Context) {
	var comment model.CommentInfo
	comment.Content = c.Query("content")
	comment.Email = c.Query("email")
	replyid := c.Query("replyid")
	comment.ReplyArticalId, _ = strconv.Atoi(replyid)
	model.SaveComment(comment)
} //对文章进行评论，需要指明ID，同时将该评论附上ID,且留下邮箱

func JundgeEmail(email string) {
	result, _ := regexp.MatchString(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, email)
	if result {
		//邮箱正确
	} else {
		//邮箱不正确
	}
} //判断邮箱是否存在

func ReplyComment(c *gin.Context) {
	var reply model.ReplyInfo
	reply.Content = c.Query("content")
	id := c.Query("replycommentid")
	reply.ReplyCommentId, _ = strconv.Atoi(id)
	email := model.FindCommentEmail(reply.ReplyCommentId)
	model.SaveReply(reply)
	go SendEamil(email)
} //对评论进行回复，需要指明评论的ID，同时发送邮件告诉对方有人回复

func GetComment(c *gin.Context) {
	ID := c.Query("id") //ID为文章ID
	Id, _ := strconv.Atoi(ID)
	comments := model.FindComment(Id)
	c.JSON(http.StatusOK, gin.H{
		"this is the comments about": Id,
	})
	for _, a := range comments {
		fmt.Println(a.CommentID, a.Content)
		replys := model.FindReply(a.CommentID)
		for _, b := range replys {
			fmt.Println("   ", b.ReplyID, b.Content)
		}
	}
} //查询评论并同时返回回复

package controller

import (
	"Test/model"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadComment(c *gin.Context) {
	var comment model.CommentInfo
	comment.Content = c.PostForm("content")
	comment.Email = c.PostForm("email")
	replyid := c.PostForm("replyid")
	comment.ReplyArticalId, _ = strconv.Atoi(replyid)
	if Jundgeemail(comment.Email) {
		model.SaveComment(comment)
		c.JSON(http.StatusOK, gin.H{
			"content": comment.Content,
			"comment": "upload success",
		})
		//邮箱正确
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": "email is not exist",
		}) //邮箱不正确
	}
} //对文章进行评论，需要指明ID，同时将该评论附上ID,且留下邮箱

func JundgeEmail(email string) bool {
	result, _ := regexp.MatchString(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, email)
	return result
} //判断邮箱是否存在

func Jundgeemail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func ReplyComment(c *gin.Context) {
	var reply model.ReplyInfo
	reply.Content = c.PostForm("content")
	id := c.PostForm("id") //id为要回复的评论的id
	reply.ReplyCommentId, _ = strconv.Atoi(id)
	email := model.FindCommentEmail(reply.ReplyCommentId)
	model.SaveReply(reply)
	go SendEamil(email)
	c.JSON(http.StatusOK, gin.H{
		"replycontent": reply.Content,
		"upload":       "success",
	})
} //对评论进行回复，需要指明评论的ID，同时发送邮件告诉对方有人回复

func GetComment(c *gin.Context) {
	ID := c.Query("id") //ID为文章ID
	Id, _ := strconv.Atoi(ID)
	comments := model.FindComment(Id)
	c.JSON(http.StatusOK, model.ApiReturn(comments.Msg, comments.Data))
} //查询评论并同时返回回复

func GetCommentByID(c *gin.Context) {
	ID := c.Query("id")
	Id, _ := strconv.Atoi(ID)
	comment := model.FindCommentByID(Id)
	c.JSON(http.StatusOK, model.ApiReturn(comment.Msg, comment.Data))
}

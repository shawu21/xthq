package model

import "github.com/gin-gonic/gin"

type Article struct {
	ID       int
	Content  string
	Category string
	Tag      string
}

type Comment struct {
	CommentID      int
	ReplyArticalID int
	Content        string
	Email          string
}

type Reply struct {
	ReplyCommentID int
	Content        string
}

func SaveArtical(c *gin.Context{

})//在数据库里保存文章

func SaveComment(c *gin.Context{

})//在数据库里保存评论

func SaveReply(c *gin.Context{

})//在数据库里保存回复

func FindArtical(ID int,c *gin.Context{

})查找数据库里对应ID的文章

func FindComment(ID int,c *gin.Context{

})//查找数据里对应ID的评论


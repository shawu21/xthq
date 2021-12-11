package model

import "github.com/gin-gonic/gin"

type ReturnType struct {
	Msg  string
	Data interface{}
}

func ApiReturn(msg string, data interface{}) gin.H {
	return gin.H{
		"message": msg,
		"data":    data,
	}
}

func FindComment(ID int) ReturnType {
	var Comments []CommentInfo
	err := db.Where("reply_artical_id = ?", ID).Find(&Comments).Error
	if err != nil {
		return ReturnType{Msg: "查询失败", Data: gin.H{
			"error": err,
		}}
	} else {
		return ReturnType{Msg: "查询成功", Data: gin.H{
			"data": Comments,
		}}
	}
} //查找数据里对应ID的评论

func FindCommentByID(ID int) ReturnType {
	var Comment CommentInfo
	db.Where("comment_id = ?", ID).Find(&Comment)
	rep := FindReply(Comment.CommentID)
	return ReturnType{Msg: "查询成功", Data: gin.H{
		"comment": Comment.Content,
		"reply":   rep.Data,
	}}
}

func FindCommentEmail(ID int) string {
	var comment CommentInfo
	db.Where("comment_id = ?", ID).Find(&comment)
	return comment.Email
}

func SaveComment(Comment CommentInfo) {
	db.AutoMigrate(&CommentInfo{})
	db.Create(&Comment)
} //在数据库里保存评论

func FindReply(ID int) ReturnType {
	var replys []ReplyInfo
	db.Where("reply_comment_id = ?", ID).Find(&replys)
	return ReturnType{Msg: "查询成功", Data: gin.H{
		"data": replys,
	}}
} //返回对应commentid的回复切片

func SaveReply(Reply ReplyInfo) {
	db.AutoMigrate(&ReplyInfo{})
	db.Create(&Reply)
} //在数据库里保存回复

package model

import "github.com/gin-gonic/gin"

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
	err := db.Where("comment_id = ?", ID).Find(&Comment).Error
	rep := FindReply(Comment.CommentID)
	if err != nil {
		return ReturnType{Msg: "查询失败", Data: gin.H{
			"error": err,
		}}
	} else {
		return ReturnType{Msg: "查询成功", Data: gin.H{
			"comment": Comment.Content,
			"reply":   rep.Data,
		}}
	}
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

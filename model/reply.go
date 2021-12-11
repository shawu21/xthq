package model

import "github.com/gin-gonic/gin"

type ReplyInfo struct {
	ReplyID        int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	ReplyCommentId int
	Content        string
}

func SaveReply(Reply ReplyInfo) {
	db.AutoMigrate(&ReplyInfo{})
	db.Create(&Reply)
} //在数据库里保存回复

func FindReply(ID int) ReturnType {
	var replys []ReplyInfo
	db.Where("reply_comment_id = ?", ID).Find(&replys)
	return ReturnType{Msg: "查询成功", Data: gin.H{
		"data": replys,
	}}
} //返回对应commentid的回复切片

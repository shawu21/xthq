package model

import (
	"fmt"
)

type ArticleInfo struct {
	ID       int `gorm:"AUTO_INCREMENT"`
	Content  string
	Category string
	Tag      string
}

type CommentInfo struct {
	CommentID      int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	ReplyArticalId int
	Content        string
	Email          string
}

type ReplyInfo struct {
	ReplyID        int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	ReplyCommentId int
	Content        string
}

func SaveArtical(Article ArticleInfo) {
	db.AutoMigrate(&ArticleInfo{})
	err := db.Create(&Article)
	if err != nil {
		fmt.Println("error=", err)
	}
} //在数据库里保存文章

func SaveComment(Comment CommentInfo) {
	db.AutoMigrate(&CommentInfo{})
	err := db.Create(&Comment)
	if err != nil {
		fmt.Println("error=", err)
	}
} //在数据库里保存评论

func SaveReply(Reply ReplyInfo) {
	db.AutoMigrate(&ReplyInfo{})
	err := db.Create(&Reply)
	if err != nil {
		fmt.Println("error=", err)
	}
} //在数据库里保存回复

func FindArtical(ID int) {
	var ArtiCal ArticleInfo
	db.Where("ID = ?", ID).Find(&ArtiCal)
	fmt.Println(ArtiCal.Content)
} //查找数据库里对应ID的文章

func FindComment(ID int) {
	var ComMent []CommentInfo
	db.Where("ReplyArticalId = ?", ID).Find(&ComMent)
	fmt.Println(ComMent)
} //查找数据里对应ID的评论

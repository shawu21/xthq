package model

import (
	"fmt"
)

type ArticleInfo struct {
	ID       int
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
	db.Create(&Article)
} //在数据库里保存文章

func FindArtical(ID int) string {
	var ArtiCal ArticleInfo
	db.Where("ID = ?", ID).Find(&ArtiCal)
	return ArtiCal.Content
} //查找数据库里对应ID的文章

func FindArticalIDbyCate(cate string) {
	var ArtiCal []ArticleInfo
	db.Where("Category = ?", cate).Find(&ArtiCal)
	for _, a := range ArtiCal {
		fmt.Println(a.ID)
	}
} //查找数据库里对应category的文章，输出ID

func FindArticalIDbytag(tag string) {
	var ArtiCal []ArticleInfo
	db.Where("Tag = ?", tag).Find(&ArtiCal)
	for _, a := range ArtiCal {
		fmt.Println(a.ID)
	}
} //查找数据库里对应tag的文章，输出ID

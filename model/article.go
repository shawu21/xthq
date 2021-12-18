package model

import "github.com/gin-gonic/gin"

type ArticleInfo struct {
	ID       int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
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

func SaveArtical(Article ArticleInfo) {
	db.AutoMigrate(&Article)
	db.Create(&Article)
} //在数据库里保存文章

func FindArtical(ID int) ReturnType {
	var ArtiCal ArticleInfo
	err := db.Where("ID = ?", ID).Find(&ArtiCal).Error
	if err != nil {
		return ReturnType{Msg: "查询失败", Data: gin.H{
			"error": err,
		}}
	} else {
		return ReturnType{Msg: "查询成功", Data: gin.H{
			"article": ArtiCal,
		}}
	}
} //查找数据库里对应ID的文章

func FindArticalIDbyCate(cate string) ReturnType {
	var ArtiCals []ArticleInfo
	err := db.Where("category = ?", cate).Find(&ArtiCals).Error
	if err != nil {
		return ReturnType{Msg: "查询失败", Data: gin.H{
			"error": err,
		}}
	} else {
		return ReturnType{Msg: "查询成功", Data: gin.H{
			"articals": ArtiCals,
		}}
	}
} //查找数据库里对应category的文章，输出ID

func FindArticalIDbytag(tag string) ReturnType {
	var ArtiCals []ArticleInfo
	err := db.Where("tag LIKE ?", "%"+tag+"%").Find(&ArtiCals).Error
	if err != nil {
		return ReturnType{Msg: "查询失败", Data: gin.H{
			"error": err,
		}}
	} else {
		return ReturnType{Msg: "查询成功", Data: gin.H{
			"articals": ArtiCals,
		}}
	}
} //查找数据库里对应tag的文章，输出ID

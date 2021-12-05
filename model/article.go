package model

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

type ReplyInfo struct {
	ReplyID        int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	ReplyCommentId int
	Content        string
}

func SaveArtical(Article ArticleInfo) {
	db.AutoMigrate(&Article)
	db.Create(&Article)
} //在数据库里保存文章

func FindArtical(ID int) string {
	var ArtiCal ArticleInfo
	db.Where("ID = ?", ID).Find(&ArtiCal)
	return ArtiCal.Content
} //查找数据库里对应ID的文章

func FindArticalIDbyCate(cate string) []ArticleInfo {
	var ArtiCals []ArticleInfo
	db.Where("Category = ?", cate).Find(&ArtiCals)
	return ArtiCals
} //查找数据库里对应category的文章，输出ID

func FindArticalIDbytag(tag string) []ArticleInfo {
	var ArtiCals []ArticleInfo
	db.Where("Tag = ?", tag).Find(&ArtiCals)
	return ArtiCals
} //查找数据库里对应tag的文章，输出ID

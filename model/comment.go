package model

func FindComment(ID int) string {
	var ComMent CommentInfo
	db.Where("ReplyArticalId = ?", ID).Find(&ComMent)
	return ComMent.Content
} //查找数据里对应ID的评论

func SaveComment(Comment CommentInfo) {
	db.AutoMigrate(&CommentInfo{})
	db.Create(&Comment)
} //在数据库里保存评论

func SaveReply(Reply ReplyInfo) {
	db.AutoMigrate(&ReplyInfo{})
	db.Create(&Reply)
} //在数据库里保存回复

package model

func FindComment(ID int) []CommentInfo {
	var ComMents []CommentInfo
	db.Where("ReplyArticalId = ?", ID).Find(&ComMents)
	return ComMents
} //查找数据里对应ID的评论

func FindCommentEmail(ID int) string {
	var comment CommentInfo
	db.Where("CommentID = ?", ID).Find(&comment)
	return comment.Email
}

func SaveComment(Comment CommentInfo) {
	db.AutoMigrate(&CommentInfo{})
	db.Create(&Comment)
} //在数据库里保存评论

func FindReply(ID int) []ReplyInfo {
	var replys []ReplyInfo
	db.Where("ReplyCommentId = ?", ID).Find(&replys)
	return replys
} //返回对应commentid的回复切片

func SaveReply(Reply ReplyInfo) {
	db.AutoMigrate(&ReplyInfo{})
	db.Create(&Reply)
} //在数据库里保存回复

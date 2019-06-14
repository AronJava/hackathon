package comment

import "hackathon/utils"

// Reply 回复评论表
type Reply struct {
	ID         int `xorm:"'id' notnull pk autoincr"`
	Comment    string
	EventID    int
	ReplyID    int
	ReplyType  int
	FromUID    int
	ToUID      int
	CreateTime int `xorm:"created"`
}

// Add 添加评论
func Add(comment string, eventID, replyID, replyType, fromUID, toUID int) {
	reply := new(Reply)
	reply.Comment = comment
	reply.EventID = eventID
	reply.ReplyID = replyID
	reply.ReplyType = replyType
	reply.FromUID = fromUID
	reply.ToUID = toUID
	utils.Engine.Insert(reply)
}

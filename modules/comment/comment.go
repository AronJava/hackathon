package comment

import (
	"hackathon/utils"
	"log"
)

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

// GetCommentsByEventID 通过活动ID获取相关评论
func GetCommentsByEventID(eventID int) ([]Reply, error) {
	reply := make([]Reply, 0)
	err := utils.Engine.Where("event_id = ?", eventID).Find(&reply)
	log.Printf("comment GetCommentsByEventID. err: %v", err)
	return reply, err
}

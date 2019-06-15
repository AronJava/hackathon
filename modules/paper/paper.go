package paper

import (
	"hackathon/utils"
	"log"
)

// Event 活动表
type Event struct {
	ID              int `xorm:"'id' notnull pk autoincr"`
	Title           string
	Content         string
	UID             int `xorm:"uid"`
	Type            int
	Img             string
	EventStartTime  int
	EventEndTime    int
	SignupStartTime int
	SignupEndTime   int
	RecommendNum    int
	BrowseNum       int
	CreateTime      int `xorm:"created"`
}

// Show 展示数据
func Show(eventType int) []Event {
	event := make([]Event, 0)
	if eventType == 0 {
		utils.Engine.Find(&event)
		return event
	}
	utils.Engine.Where("type = ?", eventType).Find(&event)
	return event

}

// Add 添加活动
func Add(userID int, title, content string, eventType int, img string, eventStartTime, eventEndTime, signupStartTime, signupEndTime int) {
	event := new(Event)
	event.UID = userID
	event.Title = title
	event.Content = content
	event.Type = eventType
	event.Img = img
	event.EventStartTime = eventStartTime
	event.EventEndTime = eventEndTime
	event.SignupStartTime = signupStartTime
	event.SignupEndTime = signupEndTime
	log.Println(89779)

	affected, err := utils.Engine.Insert(event)
	log.Printf("event. affected:%v, err:%v", affected, err)

}

// Update 添加活动
func Update(userID int, title, content string, eventType int, img string, eventStartTime, eventEndTime, signupStartTime, signupEndTime, browseNum, recommendNum int) {
	event := new(Event)
	event.UID = userID
	event.Title = title
	event.Content = content
	event.Type = eventType
	event.Img = img
	event.EventStartTime = eventStartTime
	event.EventEndTime = eventEndTime
	event.SignupStartTime = signupStartTime
	event.SignupEndTime = signupEndTime
	event.BrowseNum = browseNum
	event.RecommendNum = recommendNum

	utils.Engine.Insert(event)
}

// GetInfoByEventID 通过帖子ID获取信息
func GetInfoByEventID(eventID int) (*Event, error) {
	event := new(Event)
	_, err := utils.Engine.Where("id = ?", eventID).Get(event)
	log.Printf("paper GetInfoByEventID. err: %v", err)
	return event, err

}

// GetInfoByType 通过帖子类型获取信息
func GetInfoByType(eventType int) ([]Event, error) {
	events := make([]Event, 0)
	err := utils.Engine.Where("type = ?", eventType).Desc("recommend_num").Find(&events)
	log.Printf("paper GetInfoByEventID. err: %v", err)
	return events, err

}

//根据id获取event
func GetPaperByid(id int) Event {
	var event Event
	utils.Engine.Where("id = ?", id).Get(&event)
	return event
}

//根据uid获取信息
func GetPapersByUid(id int) []Event {
	events := make([]Event, 0)
	utils.Engine.Where("uid = ?", id).Find(&events)
	return events
}

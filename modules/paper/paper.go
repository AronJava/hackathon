package paper

import "hackathon/utils"

// Event 活动表
type Event struct {
	ID              int `xorm:"'id' notnull pk autoincr"`
	Title           string
	Content         string
	UID             int `json:"uid"`
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
		utils.Engine.Get(&event)
		return event
	}
	utils.Engine.Where("type = ?", eventType).Get(&event)
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

	utils.Engine.Insert(event)
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

package paper

import "hackathon/utils"

// Event 活动表
type Event struct {
	ID              int `xorm:"'id' notnull pk autoincr"`
	Title           string
	Content         string
	UID             int `json:"uid"`
	Type            string
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
func Show() []Event {
	event := make([]Event, 0)
	utils.Engine.Get(&event)
	return event
}

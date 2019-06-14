package signup

import (
	"hackathon/utils"
)

// Signup 签到表
type Signup struct {
	ID         int `xorm:"'id' notnull pk autoincr"`
	EventID    int
	Status     int
	UserID     int
	CreateTime int `xorm:"created"`
}

// CheckIn 签到
func CheckIn(userID, eventID int) {
	signup := new(Signup)
	signup.UserID = userID
	signup.EventID = eventID
	signup.Status = 2
	utils.Engine.Insert(signup)
}

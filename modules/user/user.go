package user

import "hackathon/utils"

// User 用户信息表
type User struct {
	ID            int `xorm:"'id' notnull pk autoincr"`
	Name          string
	UserName      string
	PassWord      string
	CreatePaperID string
	StorePaperID  string
	CreateTime    int `xorm:"created"`
}

// GetInfo 获取用户信息
func GetInfo(userID int) *User {
	user := new(User)
	utils.Engine.Get(&user)
	return user
}

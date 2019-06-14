package user

import (
	"hackathon/utils"
	"log"
)

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
func GetInfo(userID int) (*User, error) {
	user := new(User)
	_, err := utils.Engine.Where("id = ?", userID).Get(&user)
	log.Printf("user GetInfo. err: %v", err)
	return user, err
}

// GetInfoByUsername 通过名字获取用户信息
func GetInfoByUsername(username string) (*User, error) {
	user := new(User)
	_, err := utils.Engine.Where("user_name = ?", username).Get(&user)
	log.Printf("user GetInfoByUsername. err: %v", err)
	return user, err
}

func Insert(u user) error {
	_, err := utils.Engine.Insert(user)
	return err
}

package admin

import (
	"net/http"
	"strconv"

	"hackathon/utils"
)

type User struct {
	Id            int    `xorm:"notnull pk autoincr"`
	Name          string
	UserName      string
	PassWord      string
	CreatePaperId string
	StorePaperId  string
	Img           string
	CreateTime    int64  `xorm:"created"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	id, _ := strconv(r.Header.Get("User-ID")[0])
	userInfo := make([]User, 0)
	err := utils.Engine.Find(&userInfo, User{Id: id})
	if err != nil {
		log.Printf("find userinfo failed. err:%v", err)
		utils.ResponseJson(w, -1, "find userInfo failed.", nil)
		return
	}
	utils.ResponseJson(w, 0, "success", userInfo)
}

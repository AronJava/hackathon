package account

import (
	"net/http"
	"log"

	"hackathon/modules/user"
	"hackathon/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	userName := vars.Get("username")
	if userName == "" {
		log.Printf("userName is invaild")
		utils.ResponseJson(w, -1, "user name is invaild", nil)
		return
	}

	passWord := vars.Get("password")
	if passWord == "" {
		log.Printf("password is invaild")
		utils.ResponseJson(w, -1, "password is invaild", nil)
		return
	}

	name := vars.Get("name")
	if userName == "" {
		log.Printf("name is invaild")
		utils.ResponseJson(w, -1, "name is invaild", nil)
		return
	}

	userInfo := user.User{
		Name:     name,
		UserName: userName,
		Password: passWord,
	}

	err := user.Insert(userInfo)
	if err != nil {
		utils.ResponseJson(w, -2, "insert error", nil)
		return
	}
	utils.ResponseJson(w, 0, "register success", nil)
}

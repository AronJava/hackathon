package account

import(
	"net/http"
	"log"
	"strconv"

	"hackathon/modules/user"
	"hackathon/modules/session"
	"hackathon/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	userName := vars.Get("userName")
	if userName == "" {
		log.Printf("userName is invaild")
		utils.ResponseJson(w, -1, "user name is invaild", nil)
		return
	}

	passWord := vars.Get("passWord")
	if passWord == "" {
		log.Printf("password is invaild")
		utils.ResponseJson(w, -1, "password is invaild", nil)
		return
	}

	userInfo, err := user.GetInfoByUsername(userName)
	if err != nil {
		log.Printf("get info failed. err:%v", err)
		utils.ResponseJson(w, -1, "get info failed.", nil)
		return
	}
	if userInfo.ID == 0 {
		log.Printf("user is not exist, userName:%v", userName)
		utils.ResponseJson(w, -2, "user is not exist", nil)
		return
	}

	if userInfo.PassWord != passWord {
		utils.ResponseJson(w, -3, "password is wrong", nil)
		return
	}

	s, _ := session.Store.Get(r, "session_name")
	s.Values["userid"] = strconv.Itoa(userInfo.ID)

	http.Redirect(w, r, "/index/show", 302)
}

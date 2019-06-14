package signup

import (
	"hackathon/modules/signup"
	"hackathon/utils"
	"log"
	"net/http"
	"strconv"
)

// CheckIn 签到
func CheckIn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userID, err := strconv.Atoi(r.Header.Get("User-Id"))
	if userID <= 0 {
		utils.ResponseJson(w, -1, "获取用户登录状态失败", nil)
		log.Printf("update get userID failed. err:%s", err)
		return
	}
	eventID, err := strconv.Atoi(r.Form.Get("eventID"))
	if err != nil {
		utils.ResponseJson(w, -1, "param error", nil)
		return
	}
	signup.CheckIn(userID, eventID)
	utils.ResponseJson(w, 0, "签到成功", nil)

}

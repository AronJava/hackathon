package signup

import (
	"hackathon/modules/signup"
	"hackathon/utils"
	"net/http"
	"strconv"
)

// CheckIn 签到
func CheckIn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userID := 4
	eventID, err := strconv.Atoi(r.Form.Get("eventID"))
	if err != nil {
		utils.ResponseJson(w, -1, "param error", nil)
		return
	}
	signup.CheckIn(userID, eventID)
	utils.ResponseJson(w, 0, "签到成功", nil)

}

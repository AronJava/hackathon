package admin

import (
	"net/http"
	"strconv"

	"hackathon/utils"
	"hackathon/modules/user"
)

func Index(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.Header["User-Id"][0])
	userInfo, _ := user.GetInfo(id)

	utils.ResponseJson(w, 0, "success", userInfo)
}

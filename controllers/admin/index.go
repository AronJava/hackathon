package admin

import (
	"net/http"
	"strconv"

	"hackathon/utils"
	"hackathon/modules/user"
)

func Index(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	id, _ := strconv(r.Header.Get("User-ID")[0])
	userInfo := user.GetUserInfo(id)

	utils.ResponseJson(w, 0, "success", userInfo)
}

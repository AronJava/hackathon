package admin

import (
	"net/http"

	"hackathon/utils"
	"hackathon/modules/user"
)

func Index(w http.ResponseWriter, r *http.Request) {
	id := 4
	userInfo, _ := user.GetInfo(id)

	utils.ResponseJson(w, 0, "success", userInfo)
}

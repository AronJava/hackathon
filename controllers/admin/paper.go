package admin

import(
	"net/http"
	"strconv"

	"hackathon/modules/paper"
	"hackathon/utils"
)

func Paper(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.Header["User-Id"][0])
	papers := paper.GetPapersByUid(id)
	utils.ResponseJson(w, 0, "get success", papers)
}

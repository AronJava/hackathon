package admin

import(
	"net/http"

	"hackathon/modules/paper"
	"hackathon/utils"
)

func Paper(w http.ResponseWriter, r *http.Request) {
	id := 4
	papers := paper.GetPapersByUid(id)
	utils.ResponseJson(w, 0, "get success", papers)
}

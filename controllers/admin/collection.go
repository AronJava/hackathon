package admin

import(
	"net/http"
	"strconv"

	"hackathon/modules/user"
	"hackathon/modules/paper"
	"hackathon/utils"
)

func Collection(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.Header.Get("User-ID")[0])
	userInfo := user.GetUserInfo(id)

	eventsId := strings.Split(userInfo.StorePaperID, ",")
	events := make(paper.Event, 0)
	for _, e := range eventsId {
		id, _ := strconv.Atoi(e)
		event := paper.GetPaperByid(id)
		events = append(events, event)
	}

	utils.ResponseJson(w, 0, "success", events)
}

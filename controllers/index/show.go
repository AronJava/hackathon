package index

import (
	"hackathon/modules/paper"
	"hackathon/utils"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Show 展示页面
func Show(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	eventType, err := strconv.Atoi(r.Form.Get("eventType"))

	if err != nil {
		utils.ResponseJson(w, -1, "param error", nil)
		return
	}
	events := paper.Show(eventType)
	if len(events) == 0 {
		utils.ResponseJson(w, 0, "数据为空", nil)
		return
	}
	if len(events) <= 10 {
		utils.ResponseJson(w, 0, "数据不足", events)
	}
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]paper.Event, 0)
	for len(result) < 10 {
		num := seed.Intn(len(events))
		exist := false
		for _, v := range events {
			if v == events[num] {
				exist = true
				break
			}
		}

		if !exist {
			result = append(result, events[num])
		}

	}
	utils.ResponseJson(w, 0, "", result)
}

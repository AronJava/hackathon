package paper

import (
	"hackathon/modules/paper"
	"hackathon/utils"
	"math/rand"
	"net/http"
	"time"
)

// Show 展示页面
func Show(w http.ResponseWriter, r *http.Request) {
	events := paper.Show()
	if len(events) == 0 {
		utils.ResponseJSON(w, 0, "数据为空", nil)
		return
	}
	if len(events) <= 10 {
		utils.ResponseJSON(w, 0, "数据不足", events)
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
	utils.ResponseJSON(w, 0, "", result)
}

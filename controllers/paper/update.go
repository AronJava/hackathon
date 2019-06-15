package paper

import (
	"encoding/json"
	"hackathon/modules/paper"
	"hackathon/utils"
	"io/ioutil"
	"net/http"
)

// Update 更新帖子数据（如 浏览量，收藏量，内容等）
func Update(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userID := 4

	data, _ := ioutil.ReadAll(r.Body)
	var requestData *RequestData
	err := json.Unmarshal(data, &requestData)
	if err != nil {
		utils.ResponseJson(w, 10001, "参数解析错误", nil)
		return
	}
	paper.Update(userID, requestData.Title, requestData.Content, requestData.Type, requestData.Img, requestData.EventStartTime, requestData.EventEndTime, requestData.SignupStartTime, requestData.SignupEndTime, requestData.RecommendNum, requestData.BrowseNum)
}

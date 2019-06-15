package paper

import (
	"encoding/json"
	"hackathon/modules/paper"
	"hackathon/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// RequestData 请求数据
type RequestData struct {
	Title           string
	Content         string
	Type            int
	Img             string
	EventStartTime  int
	EventEndTime    int
	SignupStartTime int
	SignupEndTime   int
	RecommendNum    int
	BrowseNum       int
}

// Add 添加活动
func Add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userID, err := strconv.Atoi(r.Header.Get("User-Id"))
	if userID <= 0 {
		utils.ResponseJson(w, -1, "获取用户登录状态失败", nil)
		log.Printf("account Summary get userID failed. err:%s", err)
		return
	}

	data, _ := ioutil.ReadAll(r.Body)
	var requestData *RequestData
	err = json.Unmarshal(data, &requestData)
	if err != nil {
		utils.ResponseJson(w, 10001, "参数解析错误", nil)
		return
	}
	paper.Add(userID, requestData.Title, requestData.Content, requestData.Type, requestData.Img, requestData.EventStartTime, requestData.EventEndTime, requestData.SignupStartTime, requestData.SignupEndTime)

}
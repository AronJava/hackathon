package paper

import (
	"encoding/json"
	"hackathon/modules/paper"
	"hackathon/utils"
	"io/ioutil"
	"log"
	"net/http"
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
	log.Println(34333)
	r.ParseForm()
	userID := 4

	log.Println(34)
	data, _ := ioutil.ReadAll(r.Body)
	log.Println(99)
	log.Println(r.Body)
	var requestData *RequestData
	err := json.Unmarshal(data, &requestData)
	log.Println(err)
	if err != nil {
		utils.ResponseJson(w, 10001, "参数解析错误", nil)
		return
	}
	log.Println(34346)
	paper.Add(userID, requestData.Title, requestData.Content, requestData.Type, requestData.Img, requestData.EventStartTime, requestData.EventEndTime, requestData.SignupStartTime, requestData.SignupEndTime)

}

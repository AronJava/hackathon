package comment

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"hackathon/modules/comment"
	"hackathon/modules/user"
	"hackathon/utils"
)

// RequestData 请求数据
type RequestData struct {
	Comment   string
	Img       string
	EventID   int
	ReplyID   int
	ReplyType int
	FromUID   int
	ToUID     int
}

// Add 增加评论
func Add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userID := 4
	userInfo, _ := user.GetInfo(userID)
	data, _ := ioutil.ReadAll(r.Body)
	var requestData *RequestData
	err := json.Unmarshal(data, &requestData)
	if err != nil {
		utils.ResponseJson(w, 10001, "参数解析错误", nil)
		return
	}
	comment.Add(requestData.Comment, requestData.EventID, requestData.ReplyID, requestData.ReplyType, userInfo.ID, requestData.ToUID)

}

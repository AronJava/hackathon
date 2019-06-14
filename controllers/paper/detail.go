package paper

import (
	"hackathon/modules/comment"
	"hackathon/modules/paper"
	"hackathon/utils"
	"net/http"
	"strconv"
)

// Comments 返回的评论内容
type ResponseComment struct {
	ReplyID       int
	Content       string
	FromUID       int
	SecondComment comment.Reply
}

// Detail 帖子详情页
func Detail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	eventID, err := strconv.Atoi(r.Form.Get("eventID"))
	if err != nil {
		utils.ResponseJson(w, -1, "param error", nil)
		return
	}
	var data []ResponseComment
	var responseComment ResponseComment
	interMap := make(map[int]comment.Reply)
	eventInfo, _ := paper.GetInfoByEventID(eventID)
	comments, _ := comment.GetCommentsByEventID(eventID)
	for _, v := range comments {
		if v.ReplyType == 2 {
			interMap[v.ReplyID] = v
			continue
		}

	}
	for _, j := range comments {
		if j.ReplyType == 2 {
			continue
		}
		responseComment.Content = j.Content
		responseComment.FromUID = j.FromUID
		if _, ok := interMap[j.ID]; ok {
			responseComment.SecondComment = interMap[j.ID]
		}
		data = append(data, responseComment)
	}

	recommendEvents, _ := paper.GetInfoByType(eventInfo.Type)
	if len(recommendEvents) <= 3 {
		utils.ResponseJson(w, 0, "", struct {
			UID             int           `json:"uid"`
			Content         string        `json:"content"`
			RecommendEvents []paper.Event `json:"recommendEvents"`
		}{eventInfo.UID, eventInfo.Content, recommendEvents})
	}
}

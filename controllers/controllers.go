package controllers

import (
	"net/http"

	"hackathon/controllers/index"
)

// Routers 路由
var Routers = map[string]http.HandlerFunc{
	"/": index.Index,
	// websocket 弹幕接口
	"/ws": index.Ws,
	// 帖子展示
	"/paper/show": Paper.Show,
	// 添加评论
	"/paper/addComment": Paper.addComment,
	// 报名
	"/paper/signup": Paper.Signup,
	// 个人信息
	"/admin/personalinfo": Admin.PersonalInfo,
}

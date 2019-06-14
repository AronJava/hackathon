package controllers

import (
	"net/http"

	"hackathon/controllers/bulletscreen"
	"hackathon/controllers/admin"
	"hackathon/controllers/account"
)

// Routers 路由
var Routers = map[string]http.HandlerFunc{
	"/bulletscreen/index": bulletscreen.Index,
	// websocket 弹幕接口
	"/bulletscreen/ws": bulletscreen.Ws,
	// 帖子展示
	"/paper/show": Paper.Show,
	// 添加评论
	"/paper/addComment": Paper.addComment,
	// 报名
	"/paper/signup": Paper.Signup,
	// 个人信息
	"/admin/index": Admin.Index,
	//我的贴子
	"/admin/paper": admin.Paper,
	//我的收藏
	"/admin/collection": admin.Collection,
	//登陆
	"/account/login": account.Login,
	//注册
	"/account/register": account.Register,
}

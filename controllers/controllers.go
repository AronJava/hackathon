package controllers

import (
	"net/http"

	"hackathon/controllers/account"
	"hackathon/controllers/admin"
	"hackathon/controllers/bulletscreen"
	"hackathon/controllers/index"
	"hackathon/controllers/paper"
	"hackathon/controllers/signup"
)

// Routers 路由
var Routers = map[string]http.HandlerFunc{
	"/bulletscreen/index": bulletscreen.Index,
	// websocket 弹幕接口
	"/bulletscreen/ws": bulletscreen.Ws,
	// 帖子展示
	"/index/show": index.Show,
	// 帖子详情
	"/paper/detail": paper.Detail,
	// 帖子详情
	"/paper/update": paper.Update,
	// 添加评论
	"/paper/add": paper.Add,
	// 上传图片
	"/paper/uplod": paper.GenerateImg,
	// 报名
	"/signup/checkin": signup.CheckIn,
	// 个人信息
	"/admin/index": admin.Index,
	//我的贴子
	"/admin/paper": admin.Paper,
	//我的收藏
	"/admin/collection": admin.Collection,
	//登陆
	"/account/login": account.Login,
	//注册
	"/account/register": account.Register,
}

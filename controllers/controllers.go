package controllers

import (
	"net/http"

	"hackathon/controllers/account"
	"hackathon/controllers/admin"
	"hackathon/controllers/bulletscreen"
	"hackathon/controllers/comment"
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
	// 帖子更新
	"/paper/update": paper.Update,
	// 添加帖子（帖子）
	"/paper/add": paper.Add,
	// 上传图片
	"/paper/upload": paper.GenerateImg,
	// 添加评论
	"/comment/add": comment.Add,
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

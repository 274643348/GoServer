package controllers

import (
	"github.com/astaxie/beego"
	"learngo/GoServer/bokeProject/beeBokeProject/models"
)

type ReplyController struct {
	beego.Controller
}

func(this * ReplyController)Add(){
	tid := this.Input().Get("tid")
	nickName := this.Input().Get("nickname")
	content := this.Input().Get("content")

	err := models.AddReply(tid,nickName,content)

	if err != nil {
		beego.Error(err)
	}

	//通过智能路由访问
	this.Redirect("/topic/view/"+tid,302)
}

func(this * ReplyController)Delete(){
	if !checkAccount(this.Ctx) {
		return
	}
	tid := this.Input().Get("tid")

	rid := this.Input().Get("rid")
	//删除回复（根据回复id）
	err := models.DeleteReply(rid)

	if err != nil {
		beego.Error(err)
	}

	//通过智能路由重定向到（tid对应文章界面）
	this.Redirect("/topic/view/"+tid,302)
}
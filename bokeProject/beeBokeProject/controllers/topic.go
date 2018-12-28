package controllers

import (
	"github.com/astaxie/beego"
	"learngo/GoServer/bokeProject/beeBokeProject/models"
)

type TopicController struct {
	beego.Controller
}

func(this * TopicController)Get(){
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"

	topics,err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	}else {
		this.Data["Topics"] = topics
	}
}

func(this * TopicController)Post(){
	//登录前提下才能增加文章
	if !checkAccount(this.Ctx) {
		this.Redirect("/login",302)
		return
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")

	var err error
	err = models.AddTopic(title,content)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic",302)
}

//匹配自动路由中的"增加文章"
func(this * TopicController)Add(){
	this.TplName = "topic_add.html"
}


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
	tid := this.Input().Get("tid");
	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title,content)
	}else{
		err = models.ModifyTopic(tid,title,content)
	}

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic",302)
}

//匹配自动路由中的"增加文章"
func(this * TopicController)Add(){
	this.TplName = "topic_add.html"
}

//匹配自动路由中的"显示文章"
func(this * TopicController)View(){
	this.TplName = "topic_view.html"

	//"/login/view?id=12"
	//tid := this.Input().Get("id");

	//"/login/view/12"
	tid := this.Ctx.Input.Param("0");
	topic,err := models.GetTopic(tid);

	if err != nil {
		beego.Error(err.Error())
	}

	//Tid用于修改操作的凭证
	this.Data["Tid"] = tid
	this.Data["Topic"] = topic
}

//匹配自动路由中的"修改文章"
func(this * TopicController)Modify(){
	this.TplName = "topic_modify.html"

	//"/login/view?id=12"
	tid := this.Input().Get("tid");
	topic,err := models.GetTopic(tid);

	if err != nil {
		beego.Error(err.Error())
	}

	//Tid用于修改操作的凭证
	this.Data["Tid"] = tid
	this.Data["Topic"] = topic
}


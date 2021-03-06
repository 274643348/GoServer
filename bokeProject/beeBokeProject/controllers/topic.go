package controllers

import (
	"github.com/astaxie/beego"
	"learngo/GoServer/bokeProject/beeBokeProject/models"
	"path"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func(this * TopicController)Get(){
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"

	topics,err := models.GetAllTopics("","",false)
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
	category := this.Input().Get("category")
	tid := this.Input().Get("tid");
	label := this.Input().Get("lable");

	_,fh,err := this.GetFile("attachment");
	if err != nil {
		beego.Error(err)
	}

	var attachment string
	if fh !=nil {
		attachment = fh.Filename
		beego.Info("ljy------attachment:",attachment)
		err = this.SaveToFile("attachment",path.Join("attachment",attachment))
		if err != nil{
			beego.Error(err)
		}
	}

	if len(tid) == 0 {
		err = models.AddTopic(title,category,label,content,attachment)
	}else{
		err = models.ModifyTopic(tid,title,category,label,content,attachment)
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
		this.Redirect("/",302)
		return
	}

	//Tid用于修改操作的凭证
	this.Data["Tid"] = tid
	this.Data["Topic"] = topic
	this.Data["Labels"] = strings.Split(topic.Labels," ");

	replies,err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err.Error())
		return
	}
	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)


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

//匹配自动路由中的"删除文章"
func(this * TopicController)Delete(){

	//"/login/view?id=12"
	tid := this.Input().Get("tid");
	err := models.DeleteTopic(tid);

	if err != nil {
		beego.Error(err.Error())
	}
	this.Redirect("/topic",302)
	return
}


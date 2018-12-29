package controllers

import (
	"github.com/astaxie/beego"
	"learngo/GoServer/bokeProject/beeBokeProject/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true;
	this.TplName = "home.html"

	//设置导航栏右边登录状态
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	cate := this.Input().Get("cate")
	label := this.Input().Get("lable")
	topics,err := models.GetAllTopics(cate,label,true);
	if err != nil {
		beego.Error(err.Error())
	}else {
		this.Data["Topics"] = topics
	}

	categories,err := models.GetAllCtegories()
	if err != nil {
		beego.Error(err.Error())
		return
	}
	this.Data["Categories"] = categories
}

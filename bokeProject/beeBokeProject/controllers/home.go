package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"learngo/GoServer/bokeProject/beeBokeProject/models"
)

type HomeController struct {
	beego.Controller
	i18n.Locale
}

func (this *HomeController) Get() {

	//根据表单数据设置语言；
	if this.Input().Get("lang") == "zh-CN"{
		this.Lang = "zh-CN"
	}else{
		this.Lang = "en-US"
	}
	this.Data["Lang"] = this.Lang

	//通过控制器来处理
	this.Data["Hi"] = this.Tr("hi")

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

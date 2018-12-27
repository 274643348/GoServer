package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	//如果携带的数据带有exit，即为退出
	isExit := this.Input().Get("exit") == "true"
	beego.Warning("ljy-------------" + this.Input().Get("exit"))
	if isExit {
		this.Ctx.SetCookie("uname","",-1,"/")
		this.Ctx.SetCookie("pwd","",-1,"/")
		this.Redirect("/",302)
		return
	}

	this.TplName = "login.html"
}


func (this *LoginController) Post() {
	//this.TplName = "login.html"
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	fmt.Sprint(this.Input())

	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd {
		maxAge:=0
		if autoLogin {
			maxAge = 1<<31-1
		}
		this.Ctx.SetCookie("uname",uname,maxAge,"/")
		this.Ctx.SetCookie("pwd",pwd,maxAge,"/")

		beego.Warning("ljy----------------login-----------success")
	}else {
		beego.Warning("ljy----------------login-----------fail")
	}
	this.Redirect("/",302);

	//防止页面渲染
	return
}

func checkAccount(ctx *context.Context)bool{
	ck,err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck,err = ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	pwd := ck.Value

	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd {
		return true
	}
	return false

}
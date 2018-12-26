package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	//this.Data["Website"] = "beego.me"
	//this.Data["Email"] = "astaxie@gmail.com"
	//this.TplName = "index.tpl"

	this.Ctx.WriteString("appname：" + beego.AppConfig.String("appname")+
		"\nhttpPort："+ beego.AppConfig.String("httpport") +
		"\nrunmode：" + beego.AppConfig.String("runmode"))


	//日志处理
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.SetLogFuncCall(true)

	beego.Emergency("this is emergency")
	beego.Alert("this is alert")
	beego.Critical("this is critical")
	beego.Error("this is error")
	beego.Warning("this is warning")
	beego.Notice("this is notice")
	beego.Informational("this is informational")
	beego.Debug("this is debug")

	beego.SetLevel(beego.LevelWarning)

	beego.Emergency("this is emergency2")
	beego.Alert("this is alert2")
	beego.Critical("this is critical2")
	beego.Error("this is error2")
	beego.Warning("this is warning2")
	beego.Notice("this is notice2")
	beego.Informational("this is informational2")
	beego.Debug("this is debug2")

}

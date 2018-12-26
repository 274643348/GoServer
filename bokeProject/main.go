package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"learngo/GoServer/bokeProject/beeBokeProject/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get(){
	this.Ctx.WriteString("hello world")
}
func init(){
	models.RegisterDB()
}

func main() {
	//打印orm信息
	orm.Debug = true;

	//自动建表
	orm.RunSyncdb("default",false,true);

	//注册路由
	beego.Router("/",&HomeController{})
	beego.Run()
}

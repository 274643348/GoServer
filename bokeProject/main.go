package main

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get(){
	this.Ctx.WriteString("hello world")
}
func main() {
	//注册路由
	beego.Router("/",&HomeController{})
	beego.Run()
}

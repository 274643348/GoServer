package routers

import (
	"learngo/GoServer/bokeProject/beeBokeProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})

	//注册登录路由
	beego.Router("/login",&controllers.LoginController{})
}

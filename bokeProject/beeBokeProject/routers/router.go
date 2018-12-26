package routers

import (
	"learngo/GoServer/bokeProject/beeBokeProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

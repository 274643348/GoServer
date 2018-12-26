package routers

import (
	"learngo/GoServer/beeBokerProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

package routers

import (
	"learngo/GoServer/bokeProject/beeBokeProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})

	//注册登录路由
	beego.Router("/login",&controllers.LoginController{})


	//注册分类路由
	beego.Router("/category",&controllers.CategoryController{})

	//注册文章路由
	beego.Router("/topic",&controllers.TopicController{})

    //自动路由(后缀必须是Controller，访问/Topic/Add，就会去调用TopicController中的Add)
	beego.AutoRouter(&controllers.TopicController{})
}

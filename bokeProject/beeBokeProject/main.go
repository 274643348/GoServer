package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"learngo/GoServer/bokeProject/beeBokeProject/models"
	_ "learngo/GoServer/bokeProject/beeBokeProject/routers"
	"github.com/astaxie/beego"
	"os"
)

func init(){
	models.RegisterDB()
}

func main() {
	i18n.SetMessage("en-US","conf/locale_en-US.ini")
	i18n.SetMessage("zh-CN","conf/locale_zh-CN.ini")
	beego.AddFuncMap("i18n", i18n.Tr)


	//打印orm信息
	orm.Debug = true;

	//自动建表
	orm.RunSyncdb("default",false,true);

	//创建附件目录
	os.MkdirAll("attachment",os.ModePerm)



	beego.Run()
}


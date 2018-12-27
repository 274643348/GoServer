package main

import (
	"github.com/astaxie/beego/orm"
	"learngo/GoServer/bokeProject/beeBokeProject/models"
	_ "learngo/GoServer/bokeProject/beeBokeProject/routers"
	"github.com/astaxie/beego"
)

func init(){
	models.RegisterDB()
}

func main() {
	//打印orm信息
	orm.Debug = true;

	//自动建表
	orm.RunSyncdb("default",false,true);


	beego.Run()
}


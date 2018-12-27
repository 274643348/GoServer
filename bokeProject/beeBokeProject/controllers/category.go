package controllers

import (
	"github.com/astaxie/beego"
	"learngo/GoServer/bokeProject/beeBokeProject/models"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	op := this.Input().Get("op")
	switch op {
	case "add":
		{
			name := this.Input().Get("name")
			if len(name) == 0 {
				break
			}
			// 数据库的写如
			beego.Warning(name)
			err := models.AddCategory(name)
			if err != nil {
				beego.Error(err)
			}

			this.Redirect("/category", 302)
			return
		}

	case "del":
		{
			id := this.Input().Get("id")
			if len(id) == 0 {
				break
			}
			// 数据库的删除
			err := models.DelCategory(id)
			if err != nil {
				beego.Error(err)
			}

			this.Redirect("/category", 302)
			return
		}

	}

	this.Data["IsCategory"] = true
	this.TplName = "category.html"

	var err error
	this.Data["Categories"],err = models.GetAllCtegories()
	if err != nil {
		beego.Error(err)
	}

}

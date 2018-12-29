package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
)

type AttachController struct {
	beego.Controller
}

func(this * AttachController)Get(){
	//href 中含有中文
	// 当template 模版解析数据时， {{.}} 假设 被  "你好.txt"  给取代， 当在浏览器显示源码时，href=%e4%bd%a0%e5%a5%bd ，被自动编码
	//url.QueryUnescape（） 进行解码。  这样就能在后台显示对应的中文路径或者资源
	filePath,err := url.QueryUnescape(this.Ctx.Request.RequestURI[1:])
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	f,err := os.Open(filePath)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	defer f.Close()

	_,err = io.Copy(this.Ctx.ResponseWriter,f)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
}

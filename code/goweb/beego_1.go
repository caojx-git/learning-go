package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

/**
 * 实现beego.Controller中的Get方法
 */
func (this *HomeController) Get() {
	this.Ctx.WriteString("Hello World")
}

func main() {
	//路由注册
	beego.Router("/", &HomeController{})
	beego.Run()
}

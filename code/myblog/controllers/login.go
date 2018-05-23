/**
 * 登录Contoller
 */

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

/**
 * 实现Get方法
 */
func (c *LoginController) Get() {

	//如果是退出，立即清空cookie
	isExist := c.Input().Get("exit") == "true"
	if isExist {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")

		c.Redirect("/", 301)
		return
	}

	//指定登录的模板文件
	c.TplName = "login.html"
}

/**
 *登录
 */
func (c *LoginController) Post() {
	//将登录信息输出到浏览器前台
	//c.Ctx.WriteString(fmt.Sprint(c.Input()))

	//获取登录信息
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")
	autoLogin := c.Input().Get("autoLogin") == "on"

	//设置cookie
	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		c.Ctx.SetCookie("uname", uname, maxAge, "/")
		c.Ctx.SetCookie("pwd", uname, maxAge, "/")
	}
	//重定向到首页
	c.Redirect("/", 301)
	return
}

/**
 * 校验cookie
 */
func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd
}

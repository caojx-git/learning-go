package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	c.Data["IsHome"] = true

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	topics, err := models.GetAllTopics(c.Input().Get("cate"), c.Input().Get("label"), true)
	if err != nil {
		beego.Error(err.Error)
	} else {
		c.Data["Topics"] = topics
	}

	//获取所有的分类
	categories, err := models.GetAllCategorys()

	if err != nil {
		beego.Error(err)
	}
	c.Data["Categories"] = categories
}

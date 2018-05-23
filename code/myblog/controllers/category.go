/**
 * 文章分类controller
 */
package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	//获取操作
	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		//插入文章分类
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("/category", 301)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("/category", 301)
		return
	}

	c.TplName = "category.html"
	c.Data["IsCategory"] = true
	var err error
	c.Data["Categories"], err = models.GetAllCategorys()

	if err != nil {
		beego.Error(err)
	}
}

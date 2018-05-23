/**
 * 文章controller
 */
package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"path"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
	c.TplName = "topic.html"

	var err error
	c.Data["Topics"], err = models.GetAllTopics("", "", true)

	if err != nil {
		beego.Error(err)
	}
}

func (c *TopicController) Post() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	//解析表单
	tid := c.Input().Get("tid")           //文章id
	title := c.Input().Get("title")       //文章标题
	content := c.Input().Get("content")   //文章内容
	category := c.Input().Get("category") //文章分类
	label := c.Input().Get("label")       //文章标签

	//获取附件
	_, fh, err := c.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}

	var attachment string
	if fh != nil {
		//保存附件
		attachment = fh.Filename
		beego.Info(attachment)
		err = c.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}

	if len(tid) == 0 {
		err = models.AddTopic(title, category, label, content, attachment)
	} else {
		err = models.ModifyTopic(tid, category, label, title, content, attachment)
	}

	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 302)

}

/**
 * 新增文章
 */
func (c *TopicController) Add() {
	c.TplName = "topic_add.html"
}

/**
 * 查看文章
 */
func (c *TopicController) View() {
	c.TplName = "topic_view.html"
	topic, err := models.GetTopic(c.Ctx.Input.Params()["0"])
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}

	c.Data["Topic"] = topic
	c.Data["Labels"] = strings.Split(topic.Label, " ")
	c.Data["Tid"] = c.Ctx.Input.Params()["0"]

	//获取所有的评论数据
	replies, err := models.GetAllReplies(c.Ctx.Input.Params()["0"])
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Replies"] = replies
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}

/**
 * 修改文章
 */
func (c *TopicController) Modify() {
	c.TplName = "topic_modify.html"

	tid := c.Input().Get("tid")

	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}

	c.Data["Topic"] = topic
	c.Data["Tid"] = tid
}

func (c *TopicController) Delete() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	err := models.DeleteTopic(c.Ctx.Input.Params()["0"])
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/", 302)
}

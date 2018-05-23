/**
 * 文章回复
 */
package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type ReplyController struct {
	beego.Controller
}

/**
 * 添加评论
 */
func (c *ReplyController) Add() {
	tid := c.Input().Get("tid")
	err := models.AddReply(tid, c.Input().Get("nickname"), c.Input().Get("Content"))

	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic/view/"+tid, 302)
}

/**
 * 删除评论
 */
func (c *ReplyController) Delete() {
	if !checkAccount(c.Ctx) {
		return
	}

	tid := c.Input().Get("tid")
	rid := c.Input().Get("rid")
	err := models.DeleteReply(rid)

	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic/view/"+tid, 302)
}

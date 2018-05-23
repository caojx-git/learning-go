package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myblog/controllers"
	"myblog/models"
	_ "myblog/routers"
	"os"
)

//引入数据模型
func init() {
	//注册数据库
	models.RegisterDB()
}

func main() {

	//开启ORM调试模式
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)

	//注册路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})

	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})

	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")

	//创建文件目录
	os.Mkdir("attachment", os.ModePerm)

	//方式一：作为静态文件处理
	//beego.SetStaticPath("/attachment", "attachment")

	//方式一：作为Controller文件处理
	beego.Router("/attachment/:all", &controllers.AttachController{})

	//启动beego
	beego.Run()
}

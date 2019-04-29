package controllers

import (
	"beelog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	// 设置模板文件名
	this.TplName = "home.html"

	// 设置模板中的数据字段值
	this.Data["IsHome"] = true

	// 根据登录成功与否设置IsLogin的状态值
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	// 获取文章信息(按创建时间逆序), 如果有分类信息的话，显示该分类的文章
	// 没有分类信息，则显示所有文章
	topics, err := models.GetAllTopics(this.Input().Get("cate"), this.Input().Get("label"), true)
	if err != nil {
		logs.Error(err)
	} else {
		this.Data["Topics"] = topics
	}

	categories, err := models.GetAllCategories()
	if err != nil {
		logs.Error(err)
	}

	this.Data["Categories"] = categories
}

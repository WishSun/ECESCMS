package controllers

import (
	"beelog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	op := this.Input().Get("op")

	// 根据操作类型执行对应的操作
	switch op {
	case "add":
		logs.Info("进入添加！！！")
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			err = nil
		}
		// 不管添加成功与否，最后都要跳转到category页面
		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DeleteCategory(id)
		if err != nil {

		}
		this.Redirect("/category", 301)
		return
	}
	this.TplName = "category.html"
	this.Data["IsCategory"] = true

	// 设置分类信息
	this.Data["Categories"], _ = models.GetAllCategories()
	logs.Info("全部分类信息", this.Data["Categories"])
}

package controllers

import (
	"github.com/astaxie/beego"
)

type MyPasswordManagerController struct {
	beego.Controller
}

func (this *MyPasswordManagerController) Get() {
	this.TplName = "my_password_manager.html"
	this.Data["Tid"] = this.GetSession("teacher_id")
}

func (this *MyPasswordManagerController) Post() {

}

// 检查密码
func (this *MyPasswordManagerController) CheckPassword() {

}

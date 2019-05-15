package controllers

import (
	"ECESCMS/code/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}

type Message struct {
	Code int
	Url  string
}

func (this *LoginController) Get() {
	this.TplName = "login.html"

	// 在登录成功后，在session中设置登录对象
	this.SetSession("obj", this.Input().Get("obj"))

	IsAdmin := this.GetSession("obj") == "admin"
	this.Data["IsAdmin"] = IsAdmin
	if this.Input().Get("relogin") == "true" {

	}
}

// 检查登录
func (this *LoginController) CheckAccount() {
	m := &Message{}

	if this.GetSession("obj") == "admin" {
		if beego.AppConfig.String("adminNumber") == this.Input().Get("unumber") &&
			beego.AppConfig.String("adminPwd") == this.Input().Get("upwd") {

			m.Code = 0
			m.Url = "/admin_home"
			data, _ := json.Marshal(m)
			_, _ = fmt.Fprintln(this.Ctx.ResponseWriter, string(data))
			return
		}
	} else {
		if models.CheckTeacherAccount(this.Input().Get("unumber"), this.Input().Get("upwd")) {
			teacher, err := models.GetTeacher(this.Input().Get("unumber"))
			if err != nil {
				logs.Error(err)
			}
			this.SetSession("is_login", true)
			this.SetSession("teacher_name", teacher.Name)
			this.SetSession("teacher_id", teacher.Id)

			m.Code = 0
			m.Url = "/teacher_home"
			data, _ := json.Marshal(m)
			_, _ = fmt.Fprintln(this.Ctx.ResponseWriter, string(data))
			return
		}
	}

	// 登录失败
	m.Code = 1
	data, _ := json.Marshal(m)
	_, _ = fmt.Fprintln(this.Ctx.ResponseWriter, string(data))
}

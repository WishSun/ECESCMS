package controllers

import (
	"ECESCMS/code/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type TeacherManagerController struct {
	beego.Controller
}

func (this *TeacherManagerController) Get() {
	this.TplName = "teacher_manager.html"

	teachers, err := models.GetAllTeacher()
	if err != nil {
		logs.Error(err)
	}
	this.Data["Teachers"] = teachers
}

func (this *TeacherManagerController) Add() {
	number := this.Input().Get("TNumber")
	name := this.Input().Get("TName")
	pwd := this.Input().Get("TPwd")

	err := models.AddTeacher(number, name, pwd)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect("/admin/teacher_manager", 302)
}

func (this *TeacherManagerController) Delete() {
	tid := this.Input().Get("tid")

	err := models.DeleteTeacher(tid)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect("/admin/teacher_manager", 302)
}

func (this *TeacherManagerController) ResetPwd() {
	tid := this.Input().Get("tid")

	type Message struct {
		Code int
		Name string
	}
	name, err := models.ResetTeacherPwd(tid)
	var m *Message
	if err != nil {
		logs.Error(err)
		m = &Message{0, name}
	} else {
		m = &Message{0, name}
	}

	// 序列化为json后返回
	data, _ := json.Marshal(m)
	fmt.Fprintln(this.Ctx.ResponseWriter, string(data))
}

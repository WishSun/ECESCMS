package controllers

import (
	"beelog/models"
	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add() {
	tid := this.Input().Get("tid")
	err := models.AddReply(tid,
		this.Input().Get("nickname"),
		this.Input().Get("content"))
	if err != nil {
		logs.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)
}

func (this *ReplyController) Delete() {
	tid := this.Input().Get("tid")
	err := models.DeleteReply(this.Input().Get("rid"))
	if err != nil {
		logs.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)
}

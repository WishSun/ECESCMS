package controllers

import (
	"beelog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"path"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic.html"

	// 获取所有文章
	topics, err := models.GetAllTopics("", "", false)
	if err != nil {
		logs.Error(err)
	} else {
		this.Data["Topics"] = topics
	}
}

func (this *TopicController) Post() {
	// 验证是否已登录，如果未登录则重定向到登录页面
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	// 解析表单
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	tid := this.Input().Get("tid")
	category := this.Input().Get("category")
	label := this.Input().Get("label")

	// 获取附件
	_, fh, err := this.GetFile("attachment")
	if err != nil {
		logs.Error(err)
	}
	var attachment string
	if fh != nil {
		// 保存附件
		attachment = fh.Filename
		logs.Info(attachment)

		// 将附件保存到指定文件夹(这里第二个参数可以是绝对路径，也可以是相对路径)
		// 最后将附件保存在 attachment/附件文件名
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			logs.Error(err)
		}
	}

	if len(tid) == 0 { // 添加文章
		err = models.AddTopic(title, category, label, content, attachment)
	} else { // 修改文章
		err = models.ModifyTopic(tid, title, category, label, content, attachment)
	}

	if err != nil {
		logs.Error(err)
	}

	// 重定向到文章页面浏览
	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic_add.html"
}

func (this *TopicController) View() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic_view.html"

	// 这里url举例为 :  /topic/view/233/344
	// "0"就是取第一个233，"1"就是取第二个344，我们这里只取第一个，即文章的id
	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		logs.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Labels"] = strings.Split(topic.Labels, " ")

	replies, err := models.GetAllReplies(this.Ctx.Input.Param("0"))
	if err != nil {
		logs.Error(err)
		return
	}

	this.Data["Replies"] = replies
}

func (this *TopicController) Modify() {
	// 验证是否已登录
	isLogin := checkAccount(this.Ctx)
	if !isLogin {
		this.Redirect("/login", 302)
		return
	}
	this.Data["IsLogin"] = isLogin

	this.TplName = "topic_modify.html"

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		logs.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
}

func (this *TopicController) Delete() {
	// 验证是否已登录
	isLogin := checkAccount(this.Ctx)
	if !isLogin {
		this.Redirect("/login", 302)
		return
	}
	this.Data["IsLogin"] = isLogin

	err := models.DeleteTopic(this.Input().Get("tid"))
	if err != nil {
		logs.Error(err)
	}
	this.Redirect("/", 302)
}

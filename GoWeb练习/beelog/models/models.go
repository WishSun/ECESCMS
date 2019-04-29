package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 必须添加该驱动
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type Category struct {
	Id              int64 //id ， 默认是自动增长型，也是主键，类型为bigint(20)
	Title           string
	Created         time.Time `orm:"index"` // 建立索引
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64     //id ， 默认是自动增长型，也是主键，类型为bigint(20)
	Uid             int64     // 用户ID
	Title           string    // 标题
	Category        string    // 文章分类
	Labels          string    // 文章标签
	Content         string    `orm:"size(5000)"` // 博客内容（设置类型为varchar(5000)，默认为varchar(255)）
	Attachment      string    // 附件(255字节)   // string类型默认为varchar(255)
	Created         time.Time `orm:"index"` // 创建时间，类型默认为datetime，并为该字段建立索引
	Updated         time.Time `orm:"index"` // 最后更新时间
	Views           int64     `orm:"index"` // 访问数
	Author          string    //作者
	ReplyTime       time.Time // 最后回复时间
	ReplyCount      int64     // 评论数量
	ReplyLastUserId int64     //最后回复用户ID
}

// 评论
type Reply struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

// 注册数据库
func RegisterDB() {
	// 注册model
	orm.RegisterModel(new(Category), new(Topic), new(Reply))

	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:xiwang1997@/beeblog?charset=utf8")

	// 自动建表
	orm.RunSyncdb("default", false, true)
}

// 添加文章
func AddTopic(title, category, label, content, attachment string) error {
	// 处理标签
	// 空格作为多个标签的分隔符，因为label中可能有多个标签，所以先使用Split方法
	// 将label按空格分割成一个slice切片，Join方法是将这个slice切片变成一个字符串，
	// 第二个参数"#$"是分隔符，
	/*
			例如，有两个标签beego和orm，那么存到数据库中就会是"$beego#$orm#"
		    而有一个标签beego的话，则是$beego#

			这样有利于模糊搜索！
	*/
	label = "$" + strings.Join(
		strings.Split(label, " "), "#$") + "#"

	o := orm.NewOrm()

	topic := &Topic{
		Title:      title,
		Content:    content,
		Category:   category,
		Labels:     label,
		Attachment: attachment,
		Created:    time.Now(),
		Updated:    time.Now(),
		ReplyTime:  time.Now(),
	}

	_, err := o.Insert(topic)
	if err != nil {
		return err
	}

	// 更新分类统计
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		// 如果不存在，简单地忽略更新操作
		cate.TopicCount++
		_, err = o.Update(cate)
	}

	return err
}

// 添加分类
func AddCategory(name string) error {
	// 获取一个新的ORM操作句柄
	o := orm.NewOrm()

	// 创建一个Title为name的分类结构
	cate := &Category{Title: name}

	// 获取操作category表的操作句柄
	qs := o.QueryTable("category")

	// 在category表中查找分类title字段值为name的分类信息并填充到cate中
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		logs.Info("已经找到，不插入")
		return err // 已经找到，说明其存在，直接返回
	}
	logs.Info("未找到，插入新分类")
	// 如果不存在，则插入该分类信息
	cate.Created = time.Now()
	cate.TopicTime = time.Now()
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

// 添加评论
func AddReply(tid, nickName, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	reply := &Reply{
		Tid:     tidNum,
		Name:    nickName,
		Content: content,
		Created: time.Now(),
	}

	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err != nil {
		return err
	}

	// 更新文章评论数和最后回复时间
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = time.Now()
		topic.ReplyCount++
		_, err = o.Update(topic)
	}
	return err
}

// 删除分类
func DeleteCategory(id string) error {
	// 将string类型的id解析为10进制的64位数
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	// 必须根据主键ID进行删除
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

// 删除文章
func DeleteTopic(id string) error {
	// 将string类型的id解析为10进制的64位数
	tidNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	var oldCate string
	o := orm.NewOrm()

	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		_, err = o.Delete(topic)
		if err != nil {
			return err
		}
	}
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		// 在category中找到旧分类，更新旧分类的文章数-1
		err = qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}
	return err
}

// 删除评论
func DeleteReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	var tidNum int64
	reply := &Reply{Id: ridNum}
	if o.Read(reply) == nil {
		tidNum = reply.Tid
		_, err = o.Delete(reply)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	// 获取到该文章的所有评论到replies中，并按创建时间逆序排列，这样最新创建的评论就会在下标0位置
	replies := make([]*Reply, 0)
	qs := o.QueryTable("Reply")
	_, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)
	if err != nil {
		return err
	}

	logs.Info("删除后评论数为 : ", len(replies))
	topic := &Topic{}
	// 找到指定文章，更新它的最新评论时间和评论数
	qs = o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return err
	}

	// 更新评论数
	topic.ReplyCount = int64(len(replies))

	// 更新最后回复时间
	if len(replies) <= 0 {
		topic.ReplyTime = time.Now()
	} else if o.Read(topic) == nil {
		topic.ReplyTime = replies[0].Created
	}
	// 更新数据库表
	_, err = o.Update(topic)
	if err != nil {
		logs.Info(err)
	}
	return err
}

// 获取所有分类
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)

	qs := o.QueryTable("category")

	// 取得所有分类填入到cates中
	_, err := qs.All(&cates)
	return cates, err
}

// 获取所有文章
func GetAllTopics(cate string, label string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")

	var err error
	if isDesc {
		if len(cate) > 0 {
			// 如果有分类信息，则先按照分类信息过滤文章
			qs = qs.Filter("category", cate)
		}
		if len(label) > 0 {
			// 如果有标签信息，则再按照标签进行过滤文章, 如果文章的labels字段包含label内容，则被筛选出来
			qs = qs.Filter("labels__contains", "$"+label+"#")
		}
		// 查询所有文章，结果按照创建顺序逆序排列
		qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

// 获取指定文章的所有评论
func GetAllReplies(tid string) ([]*Reply, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	replies := make([]*Reply, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("Reply")
	// 获取该文章的所有回复
	_, err = qs.Filter("tid", tidNum).All(&replies)

	return replies, err
}

// 获取某篇文章
func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topic := new(Topic)

	// 根据id在topic中查询文章
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	// 更新浏览数
	topic.Views++
	_, err = o.Update(topic)

	// 显示的时候，标签还是以空格分隔形式显示
	topic.Labels = strings.Replace(strings.Replace(
		topic.Labels, "#", " ", -1), "$", "", -1)
	return topic, err
}

// 修改某篇文章
func ModifyTopic(tid, title, category, label, content, attachment string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	// 处理标签
	label = "$" + strings.Join(
		strings.Split(label, " "), "#$") + "#"

	var oldCate, oldAttach string // 旧分类, 旧附件

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	// Read方法会检测topic哪些字段被赋非0值，然后按照这些值的限定去查找匹配的记录
	if o.Read(topic) == nil {
		oldCate = topic.Category
		oldAttach = topic.Attachment
		topic.Title = title
		topic.Labels = label
		topic.Content = content
		topic.Attachment = attachment
		topic.Updated = time.Now() // 更新最后修改时间
		topic.Category = category

		_, err = o.Update(topic) // 将修改更新到数据库
		if err != nil {
			return err
		}
	}

	// 更新分类统计，旧分类文章数-1
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")

		err := qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}

	// 删除旧的附件
	if len(oldAttach) > 0 {
		os.Remove(path.Join("attachment", oldAttach))
	}

	// 新分类文章数+1
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	return err
}

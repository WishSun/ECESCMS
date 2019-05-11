package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

// 添加课程
func AddCourse(number, nameCN, nameEN, knowledgeField, usedMajors, preparatoryCourse, recommendTextBook, referenceBook, relativeWebsite string) error {
	o := orm.NewOrm()

	course := &Course{
		Number:            number,
		NameCN:            nameCN,
		NameEN:            nameEN,
		KnowledgeField:    knowledgeField,
		UsedMajors:        usedMajors,
		PreparatoryCourse: preparatoryCourse,
		RecommendTextBook: recommendTextBook,
		ReferenceBook:     referenceBook,
		RelativeWebsite:   relativeWebsite,
	}

	_, err := o.Insert(course)
	return err
}

// 获取指定课程信息
func GetCourse(cid string) (*Course, error) {
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	course := new(Course)

	// 根据id在Course中查询课程
	qs := o.QueryTable("Course")
	err = qs.Filter("id", cidNum).One(course)
	if err != nil {
		return nil, err
	}

	return course, nil
}

// 获取所有课程
func GetAllCourse() ([]*Course, error) {
	o := orm.NewOrm()

	courses := make([]*Course, 0)
	qs := o.QueryTable("Course")
	_, err := qs.All(&courses)
	return courses, err
}

// 修改课程
func ModifyCourse(cid, number, nameCN, nameEN, knowledgeField, usedMajors, preparatoryCourse, recommendTextBook, referenceBook, relativeWebsite string) error {
	// 将数值字段的string转为int64
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	course := &Course{Id: cidNum}

	// Read方法会检测Course哪些字段被赋非0值，然后按照这些值的限定去查找匹配的记录
	err = o.Read(course)
	if err == nil {
		course.Number = number
		course.NameCN = nameCN
		course.NameEN = nameEN
		course.KnowledgeField = knowledgeField
		course.UsedMajors = usedMajors
		course.PreparatoryCourse = preparatoryCourse
		course.RecommendTextBook = recommendTextBook
		course.ReferenceBook = referenceBook
		course.RelativeWebsite = relativeWebsite

		_, err = o.Update(course) // 将修改更新到数据库
		return err
	}
	return err
}

// 删除指定课程
func DeleteCourse(cid string) error {
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	course := &Course{Id: cidNum}

	if o.Read(course) == nil {
		_, err = o.Delete(course)
		return err
	}
	return err
}

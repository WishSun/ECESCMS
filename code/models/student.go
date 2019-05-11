package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

// 添加学生
func AddStudent(number, name, class, grade, majorName string) error {
	o := orm.NewOrm()

	student := &Student{
		Number:     number,
		Name:       name,
		Major_name: majorName,
		Class:      class,
		Grade:      grade,
	}

	_, err := o.Insert(student)
	return err
}

// 根据专业、年级、班级来查找学生
func GetAllStudent(majorName, grade, class string) ([]Student, error) {
	o := orm.NewOrm()

	students := make([]Student, 0)
	var err error

	// 专业、年级、班级都不为空
	_, err = o.Raw("select * from Student where major_name=? and grade=? and class=?",
		majorName, grade, class).QueryRows(&students)

	return students, err
}

// 根据姓名查找学生
func GetStudentByName(name string) ([]Student, error) {
	o := orm.NewOrm()

	students := make([]Student, 0)
	var err error

	qs := o.QueryTable("Student")
	_, err = qs.Filter("name__icontains", name).All(&students)

	return students, err
}

// 修改学生信息
func ModifyStudent(sid, number, name, class, grade, majorName string) error {
	// 将数值字段的string转为int64
	sidNum, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	student := &Student{Id: sidNum}

	err = o.Read(student)
	if err == nil {
		student.Number = number
		student.Name = name
		student.Class = class
		student.Grade = grade
		student.Major_name = majorName

		_, err = o.Update(student) // 将修改更新到数据库
		return err
	}
	return err
}

// 删除学生
func DeleteStudent(sid string) error {
	sidNum, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	student := &Student{Id: sidNum}

	if o.Read(student) == nil {
		_, err = o.Delete(student)
		return err
	}
	return err
}

package models

import (
	"ECESCMS/code/common"
	"github.com/astaxie/beego/orm"
	"strconv"
)

// 验证教师用户
func CheckTeacherAccount(unumber, upwd string) bool {
	o := orm.NewOrm()
	teacher := &Teacher{Number: unumber}

	qs := o.QueryTable("Teacher")
	err := qs.Filter("number", unumber).One(teacher)
	if err == nil {
		if teacher.Password == upwd {
			return true
		} else {
			return false
		}
	}
	return false
}

// 添加新教师
func AddTeacher(number, name, pwd string) error {
	o := orm.NewOrm()

	teacher := &Teacher{
		Number:   number,
		Name:     name,
		Password: pwd,
	}

	_, err := o.Insert(teacher)
	return err
}

// 获取所有教师
func GetAllTeacher() ([]*Teacher, error) {
	o := orm.NewOrm()

	teachers := make([]*Teacher, 0)
	qs := o.QueryTable("Teacher")
	_, err := qs.All(&teachers)
	return teachers, err
}

// 根据教师号获取教师信息
func GetTeacher(number string) (*Teacher, error) {
	o := orm.NewOrm()

	teacher := &Teacher{}
	qs := o.QueryTable("Teacher")
	err := qs.Filter("number", number).One(teacher)

	return teacher, err
}

// 获取指定教师的课程列表
func GetTeacherCourseList(tid int64) ([]*common.TSCListType, error) {
	o := orm.NewOrm()

	tsclist := make([]*common.TSCListType, 0)
	tscs := make([]*TeacherSelectCourse, 0)
	qs := o.QueryTable("TeacherSelectCourse")
	_, err := qs.Filter("teacher_id", tid).All(&tscs)
	if err != nil{
		return tsclist, err
	}

	qs = o.QueryTable("MajorMapCourse")
	qsCourse := o.QueryTable("Course")
	qsMajor := o.QueryTable("Major")

	mmc := MajorMapCourse{}
	major := Major{}
	course := Course{}
	for _, tsc := range tscs{
		err = qs.Filter("id", tsc.MMC_id).One(&mmc)
		if err != nil{
			return tsclist, err
		}
		err = qsCourse.Filter("id", mmc.Course_id).One(&course)
		if err != nil{
			return tsclist, err
		}
		err = qsMajor.Filter("id", mmc.Major_id).One(&major)
		if err != nil{
			return tsclist, err
		}
		tsclist = append(tsclist, &common.TSCListType{tsc.Id,course.Number, course.NameCN, tsc.Grade, major.Name})
	}
	return tsclist, err
}

// 删除教师
func DeleteTeacher(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	teacher := &Teacher{Id: tidNum}

	if o.Read(teacher) == nil {
		_, err = o.Delete(teacher)
		return err
	}
	return err
}

// 重置密码
func ResetTeacherPwd(tid string) (string, error) {
	// 将数值字段的string转为int64
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return "", err
	}

	o := orm.NewOrm()
	teacher := &Teacher{Id: tidNum}

	// Read方法会检测Teacher哪些字段被赋非0值，然后按照这些值的限定去查找匹配的记录
	err = o.Read(teacher)
	if err == nil {
		teacher.Password = "000000"

		_, err = o.Update(teacher) // 将修改更新到数据库
		return teacher.Name, err
	}
	return teacher.Name, err
}

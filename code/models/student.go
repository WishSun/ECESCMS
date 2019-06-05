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
	if err != nil {
		return err
	}

	// 将学生添加到他所在的年级、专业、课程的活动以及活动项的成绩
	taIds := make([]int64, 0)
	tacIds := make([]int64, 0)
	_, err = o.Raw("select id from teach_activity where TSC_id in "+
		"(select id from teacher_select_course where mmc_id in "+
		"(select id from major_map_course where major_id in "+
		"(select id from major where name=?)))", majorName).QueryRows(&taIds)

	_, err = o.Raw("select id from teach_activity_child where TA_id in "+
		"(select id from teach_activity where TSC_id in "+
		"(select id from teacher_select_course where mmc_id in "+
		"(select id from major_map_course where major_id in "+
		"(select id from major where name=?))))", majorName).QueryRows(&tacIds)

	// 获取学生自己的Id
	var studentId int64
	err = o.Raw("select id from student where number=?", number).QueryRow(&studentId)
	if err != nil {
		return err
	}

	for _, taid := range taIds {
		sJA := &StudentJoinActivity{
			TA_id:      taid,
			Student_id: studentId,
			Result:     0,
		}
		_, err = o.Insert(sJA)
		if err != nil {
			return err
		}
	}
	for _, tacid := range tacIds {
		sJAC := &StudentJoinActivityChild{
			TAC_id:     tacid,
			Student_id: studentId,
			Result:     0,
		}
		_, err = o.Insert(sJAC)
		if err != nil {
			return err
		}
	}
	return nil
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

	// 删除学生参加活动、活动项的表项
	_, err = o.Raw("delete from student_join_activity_child where student_id=?", sid).Exec()
	if err != nil {
		return err
	}
	_, err = o.Raw("delete from student_join_activity where student_id=?", sid).Exec()
	return err
}

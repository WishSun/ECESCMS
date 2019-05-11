package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

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

package models

import (
	"ECESCMS/code/common"
	"github.com/astaxie/beego/orm"
	"strings"
)

// 获取所有年级
func GetAllGrade() ([]string, error) {
	o := orm.NewOrm()

	grades := make([]string, 0)
	_, err := o.Raw("select distinct grade from student").QueryRows(&grades)

	return grades, err
}

// 获取专业的所有教师选课
func GetMajorAllSelectCourse(mid string) ([]common.TSCType, error) {
	o := orm.NewOrm()

	tscs := make([]common.TSCType, 0)

	courses := make([]Course, 0)
	_, err := o.Raw("select * from course where id in (select course_id from major_map_course where major_id=?)", mid).QueryRows(&courses)
	if err != nil {
		return tscs, err
	}

	type TeacherTempType struct {
		Number string
		Name   string
	}

	tempTeacher := TeacherTempType{}
	for i := 0; i < len(courses); i++ {
		tsc := common.TSCType{courses[i].Id, courses[i].Number, courses[i].NameCN, "", ""}
		_ = o.Raw("select number, name from teacher where id in (select teacher_id from teacher_select_course where mmc_id in (select id from major_map_course where major_id=? and course_id=?))", mid, courses[i].Id).QueryRow(&tempTeacher)
		tsc.TeacherNumber = tempTeacher.Number
		tsc.TeacherName = tempTeacher.Name
		tscs = append(tscs, tsc)
	}

	return tscs, nil
}

// 调整专业教师选课
func ModifyMajorTSC(mid, grade string, CourseIds, teacherInfos []string) error {
	o := orm.NewOrm()
	var mmcId int64
	var tId int64
	var err error

	for i := 0; i < len(CourseIds); i++ {
		err = o.Raw("select id from major_map_course where major_id=? and course_id=?", mid, CourseIds[i]).QueryRow(&mmcId)
		if err != nil {
			return err
		}

		number := strings.Split(teacherInfos[i], "-")[0]
		err = o.Raw("select id from teacher where number=?", number).QueryRow(&tId)
		if err != nil {
			return err
		}

		// 删除之前该教师选课记录
		_, _ = o.Raw("delete from teacher_select_course where mmc_id=? and teacher_id=? and grade=?", mmcId, tId, grade).Exec()

		// 添加新的指标点
		tsc := &TeacherSelectCourse{
			MMC_id:     mmcId,
			Teacher_id: tId,
			Grade:      grade,
		}
		_, err = o.Insert(tsc)
		if err != nil{
			return err
		}
	}
	return nil
}

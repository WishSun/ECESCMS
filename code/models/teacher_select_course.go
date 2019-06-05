package models

import (
	"ECESCMS/code/common"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

// 获取所有年级
func GetAllGrade() ([]string, error) {
	o := orm.NewOrm()

	grades := make([]string, 0)
	_, err := o.Raw("select distinct grade from student").QueryRows(&grades)

	return grades, err
}

// 根据id获取选课课程基本信息
func GetTSCBaseById(tscid string) (*TeacherSelectCourse, error) {
	o := orm.NewOrm()
	tsc := &TeacherSelectCourse{}
	tscidNum, err := strconv.ParseInt(tscid, 10, 64)
	if err != nil {
		return tsc, err
	}

	err = o.Raw("select * from teacher_select_course where id=?", tscidNum).QueryRow(&tsc)
	return tsc, err
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
		if err != nil {
			return err
		}
	}
	return nil
}

// 修改选课课程基本信息
func ModifyTSCBase(_Tscid, _Term, _Credit, _TestMethod, _TotalPeriod, _TheoryPeriod, _ExperimentalPeriod, _ComputerPeriod,
	_PracticePeriod, _WeekPeriod, _ContentRelationImgPath, _TeachTargetOverview, _ClassroomTeachTargetOverview,
	_ExperimentTeachTargetOverview, _CourseTask, _TeachMethod, _RelationOtherCourse, _Category string) error {
	// 将数值字段的string转为int64
	tscid, err := strconv.ParseInt(_Tscid, 10, 64)
	if err != nil {
		return err
	}
	term, err := strconv.ParseInt(_Term, 10, 64)
	if err != nil {
		return err
	}
	credit, err := strconv.ParseFloat(_Credit, 64)
	if err != nil {
		return err
	}
	totalPeriod, err := strconv.ParseInt(_TotalPeriod, 10, 64)
	if err != nil {
		return err
	}
	theoryPeriod, err := strconv.ParseInt(_TheoryPeriod, 10, 64)
	if err != nil {
		return err
	}
	experimentalPeriod, err := strconv.ParseInt(_ExperimentalPeriod, 10, 64)
	if err != nil {
		return err
	}
	computerPeriod, err := strconv.ParseInt(_ComputerPeriod, 10, 64)
	if err != nil {
		return err
	}
	practicePeriod, err := strconv.ParseInt(_PracticePeriod, 10, 64)
	if err != nil {
		return err
	}
	weekPeriod, err := strconv.ParseInt(_WeekPeriod, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	tsc := &TeacherSelectCourse{Id: tscid}

	// Read方法会检测major哪些字段被赋非0值，然后按照这些值的限定去查找匹配的记录
	err = o.Read(tsc)
	if err == nil {
		tsc.Term = term
		tsc.Credit = credit
		tsc.TestMethod = _TestMethod
		tsc.TotalPeriod = totalPeriod
		tsc.TheoryPeriod = theoryPeriod
		tsc.ExperimentalPeriod = experimentalPeriod
		tsc.ComputerPeriod = computerPeriod
		tsc.PracticePeriod = practicePeriod
		tsc.WeekPeriod = weekPeriod
		tsc.ContentRelationImgPath = _ContentRelationImgPath
		tsc.TeachTargetOverview = _TeachTargetOverview
		tsc.ClassroomTeachTargetOverview = _ClassroomTeachTargetOverview
		tsc.ExperimentTeachTargetOverview = _ExperimentTeachTargetOverview
		tsc.CourseTask = _CourseTask
		tsc.TeachMethod = _TeachMethod
		tsc.RelationOtherCourse = _RelationOtherCourse
		tsc.Category = _Category

		_, err = o.Update(tsc) // 将修改更新到数据库
		return err
	}
	return err
}

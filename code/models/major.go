package models

import (
	"ECESCMS/code/common"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
)

// 添加新专业
func AddMajor(_MajorNumber, _MajorName, _TrTOverview, _MainSubject, _Degree, _CoreCourse string,
	_StudyYears, _TotalCredits, _TheoryCredits, _PracticeCredits int64) error {
	o := orm.NewOrm()

	major := &Major{
		Number:          _MajorNumber,
		Name:            _MajorName,
		Overview:        _TrTOverview,
		MainSubject:     _MainSubject,
		Degree:          _Degree,
		CoreCourse:      _CoreCourse,
		StudyYears:      _StudyYears,
		TotalCredits:    _TotalCredits,
		TheoryCredits:   _TheoryCredits,
		PracticeCredits: _PracticeCredits,
	}

	_, err := o.Insert(major)
	return err
}

// 添加毕业要求
func AddGraduationRequirement(_GRNumber, _GRName, _GRMajor_id, _GRDescription string) error {
	major_id, err := strconv.ParseInt(_GRMajor_id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	gr := &GraduationRequirement{
		Number:      _GRNumber,
		Name:        _GRName,
		Major_id:    major_id,
		Description: _GRDescription,
	}

	_, err = o.Insert(gr)
	return err
}

// 获取所有专业
func GetAllMajor() ([]*Major, error) {
	o := orm.NewOrm()

	majors := make([]*Major, 0)
	qs := o.QueryTable("Major")
	_, err := qs.All(&majors)
	return majors, err
}

// 获取专业所有课程
func GetMajorAllCourse(mid string) ([]Course, error) {
	o := orm.NewOrm()

	major_courses := make([]Course, 0)
	_, err := o.Raw(
		fmt.Sprintf("SELECT * FROM course WHERE id IN (SELECT course_id FROM major_map_course WHERE major_id=%s)", mid)).QueryRows(&major_courses)
	return major_courses, err
}

// 根据mmcId获取专业名称
func GetMajorNameByMMCId(mmcId string)(string, error){
	o := orm.NewOrm()

	var major_name string
	_, err := o.Raw("select name from major where id in (select major_id from major_map_course where id=?)", mmcId).QueryRows(&major_name)
	return major_name, err
}

// 根据专业名获取专业Id
func GetMajorIdByMajorName(majorName string)(string, error){
	o := orm.NewOrm()
	var mid int64
	err := o.Raw("select id from major where name=?", majorName).QueryRow(&mid)
	smid := strconv.FormatInt(mid, 10)
	return smid, err
}

// 获取专业所有培养目标点以及目标概述
func GetMajorAllTrainTarget(mid string) ([]*TrainTarget, string, error) {
	midNum, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		return nil, "", err
	}
	o := orm.NewOrm()

	// 获取培养目标
	trts := make([]*TrainTarget, 0)
	qs := o.QueryTable("TrainTarget")
	_, err = qs.Filter("major_id", midNum).All(&trts)

	// 获取目标概述
	type Ret struct {
		Overview string
	}
	var r Ret
	err = o.Raw("select overview from major where id=?", midNum).QueryRow(&r)
	if err != nil {
		return nil, "", err
	}

	return trts, r.Overview, nil
}

// 获取专业的所有毕业要求基本信息
func GetMajorAllGraduationRequirement(mid string) ([]*GraduationRequirement, error) {
	midNum, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()

	// 获取毕业要求
	grs := make([]*GraduationRequirement, 0)
	qs := o.QueryTable("GraduationRequirement")
	_, err = qs.Filter("major_id", midNum).All(&grs)

	return grs, err
}

// 获取指定毕业要求的所有指标点
func GetMajorGRAllIndicatorPoint(grid string) ([]IndicatorPoint, error) {
	gridNum, err := strconv.ParseInt(grid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()

	var ips []IndicatorPoint
	_, err = o.Raw("select * from indicator_point where GR_id=?", gridNum).QueryRows(&ips)
	return ips, err
}

// 获取指定专业基本信息
func GetMajorBase(mid string) (*Major, error) {
	midNum, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	major := new(Major)

	// 根据id在Major中查询专业
	qs := o.QueryTable("Major")
	err = qs.Filter("id", midNum).One(major)
	if err != nil {
		return nil, err
	}

	return major, nil
}

// 修改专业培养目标
func ModifyMajorTrainTarget(trts []*common.TrTType, mid string) error {
	midNum, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	// 删除之前的培养目标
	_, err = o.Raw("delete from train_target where major_id=?", midNum).Exec()
	if err != nil {
		logs.Error(err)
	}

	// 添加新的培养目标
	for _, t := range trts {
		trt := &TrainTarget{
			Number:   t.TrTNumber,
			Content:  t.TrTContent,
			Major_id: midNum,
		}
		_, err = o.Insert(trt)
		if err != nil {
			return err
		}
	}
	return nil
}

// 修改专业基本信息
func ModifyMajorBase(_Mid, _MajorNumber, _MajorName, _TrTOverview, _MainSubject, _Degree, _CoreCourse,
	_StudyYears, _TotalCredits, _TheoryCredits, _PracticeCredits string) error {
	// 将数值字段的string转为int64
	midNum, err := strconv.ParseInt(_Mid, 10, 64)
	if err != nil {
		return err
	}
	studyYearsNum, err := strconv.ParseInt(_StudyYears, 10, 64)
	if err != nil {
		return err
	}
	totalCredits, err := strconv.ParseInt(_TotalCredits, 10, 64)
	if err != nil {
		return err
	}
	theoryCredits, err := strconv.ParseInt(_TheoryCredits, 10, 64)
	if err != nil {
		return err
	}
	practiceCredits, err := strconv.ParseInt(_PracticeCredits, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	major := &Major{Id: midNum}
	// Read方法会检测major哪些字段被赋非0值，然后按照这些值的限定去查找匹配的记录
	err = o.Read(major)
	if err == nil {
		major.Number = _MajorNumber
		major.Name = _MajorName
		major.Overview = _TrTOverview
		major.MainSubject = _MainSubject
		major.Degree = _Degree
		major.CoreCourse = _CoreCourse
		major.StudyYears = studyYearsNum
		major.TotalCredits = totalCredits
		major.TheoryCredits = theoryCredits
		major.PracticeCredits = practiceCredits

		_, err = o.Update(major) // 将修改更新到数据库
		return err
	}
	return err
}

// 修改毕业要求
func ModifyGraduationRequirement(_GRId, _GRNumber, _GRName, _GRDescription string) error {
	gr_id, err := strconv.ParseInt(_GRId, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	_, err = o.Raw("update graduation_requirement set number=?, name=?, description=? where id=?",
		_GRNumber, _GRName, _GRDescription, gr_id).Exec()
	return err
}

// 修改毕业要求指标点
func ModifyMajorGRIP(ips []*common.GRIPType, grid string) error {
	gridNum, err := strconv.ParseInt(grid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	// 删除之前的指标点
	_, err = o.Raw("delete from indicator_point where GR_id=?", gridNum).Exec()
	if err != nil {
		logs.Error(err)
	}

	// 添加新的指标点
	for _, t := range ips {
		ip := &IndicatorPoint{
			Number:      t.GRIPNumber,
			GR_id:       gridNum,
			Description: t.GRIPDescription,
		}
		_, err = o.Insert(ip)
		if err != nil {
			return err
		}
	}
	return nil
}

// 修改专业课程
func ModifyMajorCourses(mid string, cids []string) error {
	midNum, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	// 删除之前的专业课程
	_, err = o.Raw("delete from major_map_course where major_id=?", midNum).Exec()
	if err != nil {
		return err
	}

	// 添加新的专业课程
	for _, cid := range cids {
		logs.Info("新课程:[%s]", cid)
		cidNum, err := strconv.ParseInt(cid, 10, 64)
		if err != nil {
			return err
		}
		mmc := &MajorMapCourse{
			Major_id:  midNum,
			Course_id: cidNum,
		}
		_, err = o.Insert(mmc)
		if err != nil {
			return err
		}
		logs.Info("插入成功！")
	}
	return nil
}

// 删除指定专业
func DeleteMajor(mid string) error {
	midNum, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	major := &Major{Id: midNum}

	if o.Read(major) == nil {
		_, err = o.Delete(major)
		return err
	}
	return err
}

// 删除毕业要求
func DeleteGraduationRequirement(_GRId string) error {
	gr_id, err := strconv.ParseInt(_GRId, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	_, err = o.Raw("delete from graduation_requirement where id=?", gr_id).Exec()
	return err
}

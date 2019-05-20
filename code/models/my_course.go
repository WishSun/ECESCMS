package models

import (
	"ECESCMS/code/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

// 添加教学活动
func AddTSCTeachActivity(tscid, name, weight string, tets []*common.SendTeTWeight) error {
	o := orm.NewOrm()

	// 添加教学活动表项
	tscidNum, err := strconv.ParseInt(tscid, 10, 64)
	if err != nil {
		return err
	}
	weightNum, err := strconv.ParseInt(weight, 10, 64)
	if err != nil {
		return err
	}
	ta := &TeachActivity{
		TSC_id: tscidNum,
		Name:   name,
		Weight: weightNum,
	}
	_, err = o.Insert(ta)
	if err != nil {
		return err
	}
	// 获取教学活动id
	var taid int64
	err = o.Raw("select id from teach_activity where TSC_id=? and name=?", tscidNum, name).QueryRow(&taid)
	if err != nil {
		return err
	}

	// 添加教学活动的教学目标权重比
	for i := 0; i < len(tets); i++ {
		tetw := &TeachTargetWeightInTeachActivity{}
		var tetid int64
		if tets[i].TeTId == "" {
			tetid = 0
		} else {
			tetid, err = strconv.ParseInt(tets[i].TeTId, 10, 64)
			if err != nil {
				return err
			}
		}

		tetw.TeT_id = tetid
		tetw.Weight = tets[i].Weight
		tetw.TA_id = taid

		_, err = o.Insert(tetw)
		if err != nil {
			return err
		}
	}

	// 初始化学生教学活动的成绩
	students, err := GetStudentByTAId(strconv.FormatInt(taid, 10))
	for _, student := range students {
		sJA := &StudentJoinActivity{
			Student_id: student.Id,
			TA_id:      taid,
			Result:     0,
		}

		_, err := o.Insert(sJA)
		if err != nil {
			return err
		}
	}

	return nil
}

// 添加教学活动项
func AddTSCTeachActivityChild(taid, name, value string, tetIds []string) error {
	o := orm.NewOrm()

	// 添加活动项表项
	valueNum, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	taidNum, err := strconv.ParseInt(taid, 10, 64)
	if err != nil {
		return err
	}
	tac := &TeachActivityChild{
		TA_id: taidNum,
		Name:  name,
		Value: valueNum,
	}
	_, err = o.Insert(tac)
	if err != nil {
		return err
	}

	var tacid int64
	err = o.Raw("select id from teach_activity_child where TA_id=? and name=?", taid, name).QueryRow(&tacid)
	if err != nil {
		return err
	}
	// 添加活动项支撑教学目标表项
	for _, id := range tetIds {
		idNum, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return err
		}
		tacSuptTeT := &TACSuptTeT{
			TAC_id: tacid,
			TeT_id: idNum,
		}
		_, err = o.Insert(tacSuptTeT)
		if err != nil {
			return err
		}
	}

	// 初始化学生活动项成绩
	students, err := GetStudentByTAId(taid)

	// 初始化成绩都为0
	for _, student := range students {
		sjAC := &StudentJoinActivityChild{
			Student_id: student.Id,
			TAC_id:     tacid,
			Result:     0,
		}
		_, err = o.Insert(sjAC)
		if err != nil {
			return err
		}
	}
	return nil
}

// 获取课程所有教学活动
func GetTSCTeachActivity(tscid string) ([]*TeachActivity, error) {
	tas := make([]*TeachActivity, 0)
	var err error

	o := orm.NewOrm()

	_, err = o.Raw("select * from teach_activity where TSC_id=?", tscid).QueryRows(&tas)
	return tas, err
}

// 获取教学活动支撑的教学目标
func GetTASuptTeachTargets(taid string) ([]*TeachTarget, error) {
	o := orm.NewOrm()
	tets := make([]*TeachTarget, 0)

	_, err := o.Raw("select * from teach_target where id in (select TeT_id from teach_target_weight_in_teach_activity where TA_id=? and weight<>?)", taid, 0).QueryRows(&tets)
	return tets, err
}

// 获取教学活动带有权值教学目标
func GetTATeTWs(taid, tscid string) (string, int64, []*common.ModifyGetTeTWeight, error) {
	o := orm.NewOrm()
	mttws := make([]*common.ModifyGetTeTWeight, 0)

	type Temp struct {
		TAName   string
		TAWeight int64
	}
	var temp Temp
	err := o.Raw("select name from teach_activity where id=?", taid).QueryRow(&temp.TAName)
	err = o.Raw("select weight from teach_activity where id=?", taid).QueryRow(&temp.TAWeight)
	if err != nil {
		return temp.TAName, temp.TAWeight, mttws, err
	}

	tets, err := GetTSCTeachTarget(tscid)
	if err != nil {
		return temp.TAName, temp.TAWeight, mttws, err
	}
	for i := 0; i < len(tets); i++ {
		mttw := &common.ModifyGetTeTWeight{}
		mttw.Number = tets[i].Number
		mttw.Description = tets[i].Description
		err = o.Raw("select id from teach_target_weight_in_teach_activity where TeT_id=? and TA_id=?", tets[i].Id, taid).QueryRow(&mttw.TTWId)
		if err != nil {
			return temp.TAName, temp.TAWeight, mttws, err
		}
		err = o.Raw("select weight from teach_target_weight_in_teach_activity where TeT_id=? and TA_id=?", tets[i].Id, taid).QueryRow(&mttw.Weight)
		if err != nil {
			return temp.TAName, temp.TAWeight, mttws, err
		}
		mttws = append(mttws, mttw)
	}

	return temp.TAName, temp.TAWeight, mttws, nil
}

// 根据教学活动Id获取学生
func GetStudentByTAId(taid string) ([]*Student, error) {
	o := orm.NewOrm()

	students := make([]*Student, 0)
	var err error
	var majorName string
	var grade string
	err = o.Raw("select name from major where id in "+
		"(select major_id from major_map_course where id in "+
		"(select mmc_id from teacher_select_course where id in "+
		"(select TSC_id from teach_activity where id=?)))", taid).QueryRow(&majorName)
	if err != nil {
		return students, err
	}

	err = o.Raw("select grade from teacher_select_course where id in "+
		"(select TSC_id from teach_activity where id=?)", taid).QueryRow(&grade)
	if err != nil {
		return students, err
	}

	_, err = o.Raw("select * from student where major_name=? and grade=?", majorName, grade).QueryRows(&students)
	return students, err
}

// 教学活动所有子项成绩结构
type TAChildResults struct {
	Students       []*Student
	TAChilds       []*TeachActivityChild
	TAChildResults []*StudentJoinActivityChild
}

// 获取教学活动的所有教学活动项成绩
func GetTAChildResults(taid string) (*TAChildResults, error) {
	tacrs := &TAChildResults{}

	o := orm.NewOrm()

	var err error
	_, err = o.Raw("select * from teach_activity_child where TA_id=?", taid).QueryRows(&tacrs.TAChilds)
	if err != nil {
		return tacrs, err
	}

	tacrs.Students, err = GetStudentByTAId(taid)
	if err != nil {
		return tacrs, err
	}

	_, err = o.Raw("select * from student_join_activity_child where TAC_id in (select id from teach_activity_child where TA_id=?)", taid).QueryRows(&tacrs.TAChildResults)

	return tacrs, err
}

// 教学活动单个活动项成绩结构
type TAChildResult struct {
	Students       []*Student
	TAChildResults []*StudentJoinActivityChild
}

// 获取教学活动单个活动项成绩
func GetTAChildResult(tacid string) (*TAChildResult, error) {
	o := orm.NewOrm()

	tacr := &TAChildResult{}
	var taid string
	err := o.Raw("select TA_id from teach_activity_child where id=?", tacid).QueryRow(&taid)
	if err != nil {
		return tacr, err
	}

	tacr.Students, err = GetStudentByTAId(taid)
	_, err = o.Raw("select * from student_join_activity_child where TAC_id=?", tacid).QueryRows(&tacr.TAChildResults)
	return tacr, err
}

//学生指定教学活动的所有教学活动项的成绩结构
type STAChildResults struct {
	TAChilds        []*TeachActivityChild
	Student         *Student
	STAChildResults []*StudentJoinActivityChild
}

// 获取学生指定教学活动的所有教学活动项的成绩
func GetSTAChildResults(sid, taid string) (*STAChildResults, error) {
	stacrs := &STAChildResults{}

	o := orm.NewOrm()
	_, err := o.Raw("select * from teach_activity_child where TA_id=?", taid).QueryRows(&stacrs.TAChilds)
	if err != nil {
		return stacrs, err
	}

	err = o.Raw("select * from student where id=?", sid).QueryRow(&stacrs.Student)
	if err != nil {
		return stacrs, err
	}

	_, err = o.Raw("select * from student_join_activity_child where student_id=? and TAC_id in"+
		"(select id from teach_activity_child where TA_id=?)", sid, taid).QueryRows(&stacrs.STAChildResults)
	return stacrs, err
}

// 教学活动成绩结构
type TAResult struct {
	Students  []*Student
	TAName    *string
	TAId      *string
	TAResults []*StudentJoinActivity
}

// 获取教学活动成绩
func GetTAResult(taid string) (*TAResult, error) {
	tar := &TAResult{}
	o := orm.NewOrm()

	var err error

	tar.Students, err = GetStudentByTAId(taid)
	if err != nil {
		return tar, err
	}

	err = o.Raw("select name from teach_activity where id=?", taid).QueryRow(&tar.TAName)
	if err != nil {
		return tar, err
	}

	_, err = o.Raw("select * from student_join_activity where TA_id=?", taid).QueryRows(&tar.TAResults)
	return tar, err
}

// 修改教学活动
func ModifyTSCTeachActivity(taid, name, weight string, tetwms []*common.SendTeTWeightToModify) error {
	o := orm.NewOrm()

	// 修改教学活动表项
	taidNum, err := strconv.ParseInt(taid, 10, 64)
	if err != nil {
		return err
	}
	weightNum, err := strconv.ParseInt(weight, 10, 64)
	if err != nil {
		return err
	}
	ta := &TeachActivity{Id: taidNum}
	// Read方法会检测major哪些字段被赋非0值，然后按照这些值的限定去查找匹配的记录
	err = o.Read(ta)
	if err == nil {
		ta.Name = name
		ta.Weight = weightNum

		_, err = o.Update(ta) // 将修改更新到数据库
		if err != nil {
			return err
		}
	}

	// 修改教学活动的教学目标权重比
	for i := 0; i < len(tetwms); i++ {
		_, _ = o.Raw("update teach_target_weight_in_teach_activity set weight=? where id=?", tetwms[i].Weight, tetwms[i].TTWId).Exec()
	}
	return nil
}

// 删除教学活动
func DeleteTSCTeachActivity(taid string) error {
	o := orm.NewOrm()
	_, err := o.Raw("delete from teach_activity where id=?", taid).Exec()
	if err != nil {
		return err
	}
	_, err = o.Raw("delete from teach_target_weight_in_teach_activity where TA_id=?", taid).Exec()
	return err
}

// 根据教学目标Id获取毕业要求和指标点
func GetTSCTeachTargetGRAndIP(tetId string) ([]*GraduationRequirement, []*IndicatorPoint, error) {
	grs := make([]*GraduationRequirement, 0)
	ips := make([]*IndicatorPoint, 0)
	var err error

	o := orm.NewOrm()

	_, err = o.Raw("select * from graduation_requirement where "+
		"id in (select GR_id from indicator_point where id in (select IP_id from tet_supt_ip where TeT_id=?))", tetId).QueryRows(&grs)
	if err != nil {
		return grs, ips, err
	}

	ips, err = GetTSCTeachTargetIP(tetId)
	return grs, ips, err
}

// 根据教学目标Id获取指标点
func GetTSCTeachTargetIP(tetId string) ([]*IndicatorPoint, error) {
	ips := make([]*IndicatorPoint, 0)
	var err error

	o := orm.NewOrm()

	_, err = o.Raw("select * from indicator_point where "+
		"id in (select IP_id from tet_supt_ip where TeT_id=?)", tetId).QueryRows(&ips)
	return ips, err
}

// 获取指定教学活动中各个教学目标权重
func GetTATeachTargetWeight(taId string) ([]*common.TeTWeightType, error) {
	ttws := make([]*common.TeTWeightType, 0)

	taIdNum, err := strconv.ParseInt(taId, 10, 64)
	if err != nil {
		return ttws, err
	}

	temps := make([]TeachTargetWeightInTeachActivity, 0)
	o := orm.NewOrm()
	_, err = o.Raw("select * from teach_target_weight_in_teach_activity where TA_id=?", taIdNum).QueryRows(&temps)
	if err != nil {
		return ttws, err
	}

	for i := 0; i < len(temps); i++ {
		ttw := &common.TeTWeightType{}
		ttw.Weight = temps[i].Weight
		err = o.Raw("select number from teach_target where id=?", temps[i].TeT_id).QueryRow(&ttw.Number)
		if err != nil {
			return ttws, err
		}
		ttws = append(ttws, ttw)
	}

	return ttws, nil
}

// 根据专业名称和毕业要求编号获取毕业要求信息和对应的指标点信息
func GetGRAndIPs(majorName, grText string) (*GraduationRequirement, []IndicatorPoint, error) {
	o := orm.NewOrm()
	gr := &GraduationRequirement{}
	grNumber := strings.Split(grText, "-")[0]
	err := o.Raw("select * from graduation_requirement where number=? "+
		"and major_id in (select id from major where name=?)", grNumber, majorName).QueryRow(gr)

	ips := make([]IndicatorPoint, 0)
	qs := o.QueryTable("IndicatorPoint")
	_, err = qs.Filter("GR_id", gr.Id).All(&ips)
	return gr, ips, err
}

// 根据专业名称、毕业要求编号、指标点编号获取指标点信息
func GetIP(majorName, grText, ipText string) (*IndicatorPoint, error) {
	o := orm.NewOrm()

	grNumber := strings.Split(grText, "-")[0]
	ipNumber := strings.Split(ipText, "-")[0]

	ip := &IndicatorPoint{}
	err := o.Raw("select * from indicator_point where number=? "+
		"and GR_id in (select id from graduation_requirement where number=? "+
		"and major_id in (select id from major where name=?))", ipNumber, grNumber, majorName).QueryRow(ip)

	return ip, err
}

// 根据教师选课ID获取教学目标
func GetTSCTeachTarget(tscId string) ([]*TeachTarget, error) {
	o := orm.NewOrm()
	tets := make([]*TeachTarget, 0)

	qs := o.QueryTable("TeachTarget")
	_, err := qs.Filter("TSC_id", tscId).All(&tets)
	return tets, err
}

// 添加教学目标
func AddTeachTarget(tscid, number, description string) error {
	tscidNum, err := strconv.ParseInt(tscid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	tet := &TeachTarget{
		TSC_id:      tscidNum,
		Number:      number,
		Description: description,
	}

	_, err = o.Insert(tet)
	return err
}

// 修改教学目标
func ModifyTeachTarget(tetid, number, description string) error {
	// 将数值字段的string转为int64
	tetidNum, err := strconv.ParseInt(tetid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	tet := &TeachTarget{Id: tetidNum}

	// Read方法会检测TeachTarget哪些字段被赋非0值，然后按照这些值的限定去查找匹配的记录
	err = o.Read(tet)
	if err == nil {
		tet.Number = number
		tet.Description = description

		_, err = o.Update(tet) // 将修改更新到数据库
		return err
	}
	return err
}

// 修改教学目标支撑的指标点
func ModifyTeachtargetSuptIP(tetId string, ipids []string) error {
	tetIdNum, err := strconv.ParseInt(tetId, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	// 删除之前支撑的指标点
	_, err = o.Raw("delete from tet_supt_ip where TeT_id=?", tetIdNum).Exec()
	if err != nil {
		logs.Error(err)
	}

	// 添加新支撑的指标点
	for _, ipid := range ipids {
		ipidNum, err := strconv.ParseInt(ipid, 10, 64)
		if err != nil {
			return err
		}
		trt := &TeTSuptIP{
			TeT_id: tetIdNum,
			IP_id:  ipidNum,
		}
		_, err = o.Insert(trt)
		if err != nil {
			return err
		}
	}
	return nil
}

// 修改教学活动项成绩
func ModifyTACResult(tacid string, sIds, sResults []string) error {
	o := orm.NewOrm()

	for i, _ := range sIds {
		_, err := o.Raw("update student_join_activity_child set result=? where student_id=? and TAC_id=?",
			sResults[i], sIds[i], tacid).Exec()
		if err != nil {
			return err
		}
	}
	return nil
}

// 修改学生的指定教学活动的所有活动项成绩
func ModifyStudentTACResults(sid string, tacIds, stacResults []string) error {
	o := orm.NewOrm()

	for i, _ := range tacIds {
		_, err := o.Raw("update student_join_activity_child set result=? where student_id=? and TAC_id=?", stacResults[i], sid, tacIds[i]).Exec()
		if err != nil {
			return err
		}
	}
	return nil
}

// 删除教学目标
func DeleteTeachTarget(tetid string) error {
	// 删除教学目标
	tetidNum, err := strconv.ParseInt(tetid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	tet := &TeachTarget{Id: tetidNum}

	if o.Read(tet) == nil {
		_, err = o.Delete(tet)
		if err == nil {
			// 删除教学目标对指标点的支撑
			_, _ = o.Raw("delete from tet_supt_ip where TeT_id=?", tetidNum).Exec()

			// 删除教学活动项对教学目标的支撑
			_, _ = o.Raw("delete from tac_supt_tet where TeT_id=?", tetidNum).Exec()

			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

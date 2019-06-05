package models

import (
	"ECESCMS/code/common"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

// 添加教学活动
func AddTSCTeachActivity(tscid, name, weight string, resultWeight string, tets []*common.SendTeTWeight) error {
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
	resultWeightNum, err := strconv.ParseInt(resultWeight, 10, 64)
	if err != nil {
		return err
	}
	ta := &TeachActivity{
		TSC_id:       tscidNum,
		Name:         name,
		Weight:       weightNum,
		ResultWeight: resultWeightNum,
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

// 获取教学活动项
func GetTeachActivityChild(tacid string) (*TeachActivityChild, error) {
	o := orm.NewOrm()
	tac := &TeachActivityChild{}
	err := o.Raw("select * from teach_activity_child where id=?", tacid).QueryRow(tac)
	return tac, err
}

// 获取教学活动支撑的教学目标
func GetTASuptTeachTargets(taid string) ([]*TeachTarget, error) {
	o := orm.NewOrm()
	tets := make([]*TeachTarget, 0)

	_, err := o.Raw("select * from teach_target where id in (select TeT_id from teach_target_weight_in_teach_activity where TA_id=? and weight<>?)", taid, 0).QueryRows(&tets)
	return tets, err
}

// 获取教学活动项支撑的教学目标
func GetTACSuptTeachTargets(tacid string) ([]*TACSuptTeT, error) {
	o := orm.NewOrm()
	tacSuptTets := make([]*TACSuptTeT, 0)
	_, err := o.Raw("select * from tac_supt_tet where TAC_id=?", tacid).QueryRows(&tacSuptTets)
	return tacSuptTets, err
}

// 获取教学活动带有权值教学目标
func GetTATeTWs(taid, tscid string) (string, int64, int64, []*common.ModifyGetTeTWeight, error) {
	o := orm.NewOrm()
	mttws := make([]*common.ModifyGetTeTWeight, 0)

	type Temp struct {
		TAName         string
		TAWeight       int64
		TAResultWeight int64
	}
	var temp Temp
	err := o.Raw("select name from teach_activity where id=?", taid).QueryRow(&temp.TAName)
	err = o.Raw("select weight from teach_activity where id=?", taid).QueryRow(&temp.TAWeight)
	err = o.Raw("select result_weight from teach_activity where id=?", taid).QueryRow(&temp.TAResultWeight)
	if err != nil {
		return temp.TAName, temp.TAWeight, temp.TAResultWeight, mttws, err
	}

	tets, err := GetTSCTeachTarget(tscid)
	if err != nil {
		return temp.TAName, temp.TAWeight, temp.TAResultWeight, mttws, err
	}
	for i := 0; i < len(tets); i++ {
		mttw := &common.ModifyGetTeTWeight{}
		mttw.Number = tets[i].Number
		mttw.Description = tets[i].Description
		err = o.Raw("select id from teach_target_weight_in_teach_activity where TeT_id=? and TA_id=?", tets[i].Id, taid).QueryRow(&mttw.TTWId)
		if err != nil {
			return temp.TAName, temp.TAWeight, temp.TAResultWeight, mttws, err
		}
		err = o.Raw("select weight from teach_target_weight_in_teach_activity where TeT_id=? and TA_id=?", tets[i].Id, taid).QueryRow(&mttw.Weight)
		if err != nil {
			return temp.TAName, temp.TAWeight, temp.TAResultWeight, mttws, err
		}
		mttws = append(mttws, mttw)
	}

	return temp.TAName, temp.TAWeight, temp.TAResultWeight, mttws, nil
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

type MyTACSuptTeTType struct {
	TA_id    int64
	TeT_id   int64
	TAC_id   int64
	TAC_name string
}

type MyTeTSuptGR struct {
	GR_id       int64
	TeT_numbers string
}

type MyTASuptGR struct {
	GR_id    int64
	TA_names string
}

// 成绩分析报告数据结构
type ResultAnalysisData struct {
	TAs                  []*TeachActivity
	TeTs                 []*TeachTarget
	TACSuptTeT           []*MyTACSuptTeTType
	TeTSuptGR            []*MyTeTSuptGR
	TASuptGR             []*MyTASuptGR
	TeTInTAWeight        []*TeachTargetWeightInTeachActivity
	AverageResult        []*StuTeTAverageResultInTA
	LinkResult           []*LinkTotalResult
	MapTAWeightAndResult map[int64]float64
	IPs                  []*IndicatorPoint
	GRs                  []*GraduationRequirement
}

// 获取成绩分析数据
func GetResultAnalysisData(tscid string) (*ResultAnalysisData, error) {
	rad := &ResultAnalysisData{}
	rad.MapTAWeightAndResult = make(map[int64]float64)
	var err error

	rad.TAs, err = GetTSCTeachActivity(tscid)
	if err != nil {
		return rad, err
	}

	rad.TeTs, err = GetTSCTeachTarget(tscid)
	if err != nil {
		return rad, err
	}

	o := orm.NewOrm()
	// 首先清除老数据
	_, _ = o.Raw("delete from stu_tet_average_result_in_ta").Exec()
	_, _ = o.Raw("delete from link_total_result").Exec()

	_, err = o.Raw("select * from teach_target_weight_in_teach_activity where TA_id in "+
		"(select id from teach_activity where TSC_id=?)", tscid).QueryRows(&rad.TeTInTAWeight)
	if err != nil {
		return rad, err
	}

	// 找出课程的所有教学活动
	var TAIds []*int64
	_, err = o.Raw("select id from teach_activity where TSC_id=?", tscid).QueryRows(&TAIds)
	if err != nil {
		return rad, err
	}

	for _, taid := range TAIds {
		// 找出这个教学活动所支撑的所有教学目标
		var TeTIds []*int64
		_, err = o.Raw("select TeT_id from teach_target_weight_in_teach_activity where TA_id=? and weight<>0", taid).QueryRows(&TeTIds)
		if err != nil {
			return rad, err
		}

		// 计算学生数目
		var studentNum int64
		err = o.Raw("select count(id) from student_join_activity where TA_id=?", taid).QueryRow(&studentNum)
		if err != nil {
			return rad, err
		}

		// 教学目标总成绩
		var TeTTotalResult float64
		TeTTotalResult = 0

		for _, tetid := range TeTIds {
			tacids := make([]int64, 0)
			_, err = o.Raw("select TAC_id from tac_supt_tet where TeT_id=? and TAC_id in "+
				"(select id from teach_activity_child where TA_id=?)", tetid, taid).QueryRows(&tacids)
			for _, tacid := range tacids {
				tst := &MyTACSuptTeTType{}
				tst.TA_id = *taid
				tst.TeT_id = *tetid
				tst.TAC_id = tacid
				err = o.Raw("select name from teach_activity_child where id=?", tacid).QueryRow(&tst.TAC_name)
				rad.TACSuptTeT = append(rad.TACSuptTeT, tst)
			}
			/* 例如: 目标1在活动1中的总分为40
			       题目1(10分)   题目2(10分)  题目3(10分)
			学生1   1              2             3
			学生2   4              5             6
			学生3   3              2             1

			学生1的目标达成度 = (1+2+3)/(10+10+10) * 40
			学生2的目标达成都 = (4+5+6)/(10+10+10) * 40
			学生3的目标达成度 = (3+2+1)/(10+10+10) * 40

			目标达成度平均值为 ((1+2+3)/(10+10+10) * 40 + (4+5+6)/(10+10+10) * 40 + (3+2+1)/(10+10+10) * 40)/3

			同等于  ((40*(1+2+3+4+5+6+3+2+1))/(10+10+10))/3
			即累加查到的成绩表项的成绩，然后除以学生数，就等于项目达成度平均值
			*/
			var tetWeight int64
			// 找出这个教学目标在教学活动中占的分数
			err = o.Raw("select weight from teach_target_weight_in_teach_activity where TA_id=? and TeT_id=?", taid, tetid).QueryRow(&tetWeight)
			if err != nil {
				return rad, err
			}

			// 找出支撑这个教学目标的所有教学活动项
			var TACIds []*string
			_, err = o.Raw("select TAC_id from tac_supt_tet where TeT_id=?", tetid).QueryRows(&TACIds)
			if err != nil {
				return rad, err
			}

			// 筛选出属于这个教学活动的TAC_id
			var TACIds2 []*string
			sqlString := "select id from teach_activity_child where TA_id=" + strconv.FormatInt(*taid, 10) + " and id in ("
			var sep string
			for _, tacid := range TACIds {
				sqlString += sep + *tacid
				sep = ","
			}
			sqlString += ")"
			_, err := o.Raw(sqlString).QueryRows(&TACIds2)
			if err != nil {
				return rad, err
			}

			// 获取总成绩
			var totalResult float64
			// 然后在学生参与教学活动项表中找出这些TAC_id的成绩表项
			sqlString = "select SUM(result) from student_join_activity_child where TAC_id in ("
			sep = ""
			for _, tacid := range TACIds2 {
				sqlString += sep + *tacid
				sep = ","
			}
			sqlString += ")"
			err = o.Raw(sqlString).QueryRow(&totalResult)
			if err != nil {
				return rad, err
			}

			// 获取总分数
			var totalValue float64
			sqlString = "select SUM(value) from teach_activity_child where id in ("
			sep = ""
			for _, tacid := range TACIds2 {
				sqlString += sep + *tacid
				sep = ","
			}
			sqlString += ")"
			err = o.Raw(sqlString).QueryRow(&totalValue)
			if err != nil {
				return rad, err
			}

			AvgeResult := &StuTeTAverageResultInTA{}
			// 计算平均成绩
			AvgeResult.TeT_id = *tetid
			AvgeResult.TA_id = *taid
			AvgeResult.Result, err = strconv.ParseFloat(fmt.Sprintf("%.2f", ((float64(tetWeight)*totalResult)/totalValue)/float64(studentNum)), 64)
			TeTTotalResult += AvgeResult.Result
			_, err = o.Insert(AvgeResult)
			if err != nil {
				return rad, err
			}
		}

		ltr := &LinkTotalResult{
			TA_id:  *taid,
			Result: TeTTotalResult,
		}
		_, err = o.Insert(ltr)
		if err != nil {
			return rad, err
		}
		var mapKey int64
		err = o.Raw("select weight from teach_activity where id=?", *taid).QueryRow(&mapKey)
		if err != nil {
			return rad, err
		}
		rad.MapTAWeightAndResult[mapKey] = TeTTotalResult
	}

	_, err = o.Raw("select * from stu_tet_average_result_in_ta where TA_id in "+
		"(select id from teach_activity where TSC_id=?)", tscid).QueryRows(&rad.AverageResult)
	if err != nil {
		return rad, err
	}

	_, err = o.Raw("select * from link_total_result where TA_id in "+
		"(select id from teach_activity where TSC_id=?)", tscid).QueryRows(&rad.LinkResult)

	rad.GRs, rad.IPs, err = GetTSCAllIP(tscid)
	if err != nil {
		return rad, err
	}

	for _, gr := range rad.GRs {
		tetSuptGR := &MyTeTSuptGR{}
		tetSuptGR.GR_id = gr.Id
		tetSuptGR.TeT_numbers = ""

		numbers := make([]string, 0)

		_, err = o.Raw("select number from teach_target where id in "+
			"(select TeT_id from tet_supt_ip where IP_id in "+
			"(select id from indicator_point where GR_id=? and id in "+
			"(select distinct IP_id from tet_supt_ip where TeT_id in "+
			"(select id from teach_target where TSC_id=?)))) order by number", gr.Id, tscid).QueryRows(&numbers)
		sep := ""
		for _, number := range numbers {
			tetSuptGR.TeT_numbers += sep + number
			sep = ", "
		}

		rad.TeTSuptGR = append(rad.TeTSuptGR, tetSuptGR)
	}

	for _, gr := range rad.GRs {
		taSuptGR := &MyTASuptGR{}
		taSuptGR.GR_id = gr.Id
		taSuptGR.TA_names = ""

		names := make([]string, 0)
		_, err = o.Raw("select name from teach_activity where id in "+
			"(select TA_id from teach_activity_child where id in "+
			"(select TAC_id from tac_supt_tet where TeT_id in"+
			"(select TeT_id from tet_supt_ip where TeT_id in "+
			"(select id from teach_target where TSC_id=?) and IP_id in"+
			"(select id from indicator_point where GR_id=?))))", tscid, gr.Id).QueryRows(&names)
		sep := ""
		for _, name := range names {
			taSuptGR.TA_names += sep + name
			sep = ", "
		}
		rad.TASuptGR = append(rad.TASuptGR, taSuptGR)
	}
	return rad, err
}

// 获取课程信息
func GetCourseAndTSC(tscid string) (*Course, *TeacherSelectCourse, error) {
	course := &Course{}
	tsc, err := GetTSCBaseById(tscid)
	if err != nil {
		return course, tsc, err
	}
	o := orm.NewOrm()

	err = o.Raw("select * from Course where id in "+
		"(select course_id from major_map_course where id in "+
		"(select mmc_id from teacher_select_course where id=?))", tscid).QueryRow(course)
	return course, tsc, err
}

// 获取课程成绩分析报告信息
func GetTSCResultAnalysisReport(tscid string) (*ResultAnalysisReport, error) {
	o := orm.NewOrm()
	rar := &ResultAnalysisReport{}
	err := o.Raw("select * from result_analysis_report where TSC_id=?", tscid).QueryRow(rar)
	return rar, err
}

// 获取课程所支撑的所有毕业指标点
func GetTSCAllIP(tscid string) ([]*GraduationRequirement, []*IndicatorPoint, error) {
	o := orm.NewOrm()
	ips := make([]*IndicatorPoint, 0)
	grs := make([]*GraduationRequirement, 0)
	_, err := o.Raw("select * from indicator_point where id in "+
		"(select distinct IP_id from tet_supt_ip where TeT_id in "+
		"(select id from teach_target where TSC_id=?)) order by number", tscid).QueryRows(&ips)
	if err != nil {
		return grs, ips, err
	}

	_, err = o.Raw("select * from graduation_requirement where id in "+
		"(select distinct GR_id from indicator_point where id in "+
		"(select distinct IP_id from tet_supt_ip where TeT_id in "+
		"(select id from teach_target where TSC_id=?))) order by number", tscid).QueryRows(&grs)
	return grs, ips, err
}

// 修改教学活动
func ModifyTSCTeachActivity(taid, name, weight, resultWeight string, tetwms []*common.SendTeTWeightToModify) error {
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
	resultWeightNum, err := strconv.ParseInt(resultWeight, 10, 64)
	if err != nil {
		return err
	}
	ta := &TeachActivity{Id: taidNum}
	// Read方法会检测major哪些字段被赋非0值，然后按照这些值的限定去查找匹配的记录
	err = o.Read(ta)
	if err == nil {
		ta.Name = name
		ta.Weight = weightNum
		ta.ResultWeight = resultWeightNum
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

// 修改教学活动项
func ModifyTSCTeachActivityChild(tacid, tacname, tacvalue string, tetIds []string) error {
	o := orm.NewOrm()

	// 修改活动项信息
	_, err := o.Raw("update teach_activity_child set name=?,value=? where id=?", tacname, tacvalue, tacid).Exec()
	if err != nil {
		return err
	}

	// 删除旧有支持目标
	_, err = o.Raw("delete from tac_supt_tet where TAC_id=?", tacid).Exec()
	if err != nil {
		return err
	}

	// 添加新支持目标
	for _, tetId := range tetIds {
		_, err = o.Raw("insert into tac_supt_tet(TAC_id, TeT_id) values(?, ?)", tacid, tetId).Exec()
		if err != nil {
			return err
		}
	}
	return nil
}

// 修改课程成绩报告
func ModifyTSCResultAnalysisReport(id, pt, rept, aesn, el, it, tetim, gripim, tetfa string) error {
	o := orm.NewOrm()
	idNum, err := strconv.ParseInt(id, 10, 64)
	aesnNum, err := strconv.ParseInt(aesn, 10, 64)
	rar := &ResultAnalysisReport{Id: idNum}

	err = o.Read(rar)
	if err == nil {
		rar.ActualExamStudentNumber = aesnNum
		rar.ExamLocation = el
		rar.TeTImprovementMeasures = tetim
		rar.GRIPImprovementMeasures = gripim
		rar.InvigilationTeacher = it
		rar.ProblemTeacher = pt
		rar.ReadExamPaperTeacher = rept
		rar.TeTFinishAnalysis = tetfa

		_, err = o.Update(rar) // 将修改更新到数据库
		if err != nil {
			return err
		}
	}
	return err
}

// 删除教学活动
func DeleteTSCTeachActivity(taid string) error {
	o := orm.NewOrm()
	_, err := o.Raw("delete from teach_activity where id=?", taid).Exec()
	if err != nil {
		return err
	}

	// 删除教学活动支撑教学目标表项
	_, _ = o.Raw("delete from teach_target_weight_in_teach_activity where TA_id=?", taid).Exec()
	// 删除教学活动下面的所有活动项
	_, _ = o.Raw("delete from teach_activity_child where TA_id=?", taid).Exec()
	// 删除参加教学活动的学生的成绩
	_, _ = o.Raw("delete from student_join_activity where TA_id=?", taid).Exec()
	// 删除参加教学活动的所有教学活动的学生的成绩
	_, _ = o.Raw("delete from student_join_activity_child where student_id in "+
		"(select student_id from student_join_activity where TA_id=?)", taid).Exec()

	return nil
}

// 删除教学活动项
func DeleteTSCTeachActivityChild(tacid string) error {
	o := orm.NewOrm()
	// 删除教学活动项
	_, err := o.Raw("delete from teach_activity_child where id=?", tacid).Exec()
	if err != nil {
		return err
	}
	// 删除教学活动支撑教学目标
	_, err = o.Raw("delete from tac_supt_tet where TAC_id=?", tacid).Exec()
	if err != nil {
		return err
	}
	// 删除学生参加教学活动项
	_, err = o.Raw("delete from student_join_activity_child where TAC_id=?", tacid).Exec()
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

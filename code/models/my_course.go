package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

// 根据教学目标Id获取毕业要求和指标点
func GetTSCTeachTargetGRAndIP(tetId string) ([]*GraduationRequirement, []*IndicatorPoint, error) {
	grs := make([]*GraduationRequirement, 0)
	ips := make([]*IndicatorPoint, 0)
	var err error

	o := orm.NewOrm()

	_, err = o.Raw("select * from graduation_requirement where " +
		"id in (select GR_id from indicator_point where id in (select IP_id from tet_supt_ip where TeT_id=?))", tetId).QueryRows(&grs)
	if err != nil {
		return grs, ips, err
	}

	ips, err = GetTSCTeachTargetIP(tetId)
	return grs, ips, err
}

// 根据教学目标Id获取指标点
func GetTSCTeachTargetIP(tetId string)([]*IndicatorPoint, error){
	ips := make([]*IndicatorPoint, 0)
	var err error

	o := orm.NewOrm()

	_, err = o.Raw("select * from indicator_point where " +
		"id in (select IP_id from tet_supt_ip where TeT_id=?)", tetId).QueryRows(&ips)
	return ips, err
}

// 根据专业名称和毕业要求编号获取毕业要求信息和对应的指标点信息
func GetGRAndIPs(majorName, grText string)(*GraduationRequirement, []IndicatorPoint, error){
	o := orm.NewOrm()
	gr := &GraduationRequirement{}
	grNumber := strings.Split(grText, "-")[0]
	err := o.Raw("select * from graduation_requirement where number=? " +
		"and major_id in (select id from major where name=?)", grNumber, majorName).QueryRow(gr)

	ips := make([]IndicatorPoint, 0)
	qs := o.QueryTable("IndicatorPoint")
	_, err = qs.Filter("GR_id", gr.Id).All(&ips)
	return gr, ips, err
}

// 根据专业名称、毕业要求编号、指标点编号获取指标点信息
func GetIP(majorName, grText, ipText string)(*IndicatorPoint, error){
	o := orm.NewOrm()

	grNumber := strings.Split(grText, "-")[0]
	ipNumber := strings.Split(ipText, "-")[0]

	ip := &IndicatorPoint{}
	err := o.Raw("select * from indicator_point where number=? " +
		"and GR_id in (select id from graduation_requirement where number=? " +
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
func ModifyTeachtargetSuptIP(tetId string, ipids []string)error{
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
		if err != nil{
			return err
		}
		trt := &TeTSuptIP{
			TeT_id:tetIdNum,
			IP_id:ipidNum,
		}
		_, err = o.Insert(trt)
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

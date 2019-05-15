package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

// 根据教学目标Id获取毕业要求和指标点
func GetTSCTeachTargetGRAndIP(tetId string) ([]*GraduationRequirement, []*IndicatorPoint, error) {
	grs := make([]*GraduationRequirement, 0)
	ips := make([]*IndicatorPoint, 0)
	var err error

	o := orm.NewOrm()

	_, err = o.Raw("select * from graduation_requirement where in (select GR_id from indicator_point where id in (select IP_id from tet_supt_ip where TeT_id=?))", tetId).QueryRows(&grs)
	if err != nil {
		return grs, ips, err
	}
	_, err = o.Raw("select * from indicator_point where id in (select IP_id from tet_supt_ip where TeT_id=?)", tetId).QueryRows(&ips)
	if err != nil {
		return grs, ips, err
	}
	return grs, ips, err
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

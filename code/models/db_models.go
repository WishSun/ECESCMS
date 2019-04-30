package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 注册数据库
func RegisterDB() {
	// 注册model
	orm.RegisterModel(new(Major), new(TrainTarget), new(GraduationRequirement),
		new(IndicatorPoint), new(GRSuptTrT), new(Course),
		new(MajorMapCourse), new(Teacher), new(TeacherSelectCourse),
		new(TeachContent), new(TeachTarget), new(TeachActivity),
		new(TeachActivityChild), new(Student), new(TeTSuptIP),
		new(TACSuptTeT), new(StudentSelectCourse), new(StudentJoinActivityChild))

	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:xiwang1997@/ECESCMS?charset=utf8")

	// 自动建表
	orm.RunSyncdb("default", false, true)
}

//-----------------------------------
// 专业表
type Major struct {
	Id              int64  `orm:"column(id);auto"`             // 专业自增编号
	Number          string `orm:"column(number);unique"`       // 专业号
	Name            string `orm:"column(name)"`                // 专业名称
	Overview        string `orm:"column(overview);size(5000)"` // 专业概述
	TotalCredits    int    `orm:"column(total_credits)"`       // 总学分
	TheoryCredits   int    `orm:"column(theory_credits)"`      // 理论教学总学分
	PracticeCredits int    `orm:"column(practice_credits)"`    // 实践教学学分
}

// 培养目标表
type TrainTarget struct {
	Id       int64  `orm:"column(id);auto"`            // 培养目标自增编号
	Major_id string `orm:"column(major_id)"`           // 专业编号
	Content  string `orm:"column(content);size(5000)"` // 目标内容
}

//-----------------------------------
// 毕业要求表
type GraduationRequirement struct {
	Id          int64  `orm:"column(id);auto"`                // 毕业要求自增编号
	Major_id    string `orm:"column(major_id)"`               // 专业编号
	Name        string `orm:"column(name)"`                   // 要求名称
	Description string `orm:"column(description);size(5000)"` // 要求描述
}

// 指标点表
type IndicatorPoint struct {
	Id          int64  `orm:"column(id);auto"`                                             // 指标点自增编号
	GR_id       int64  `orm:"column(GR_id);description(GR意为GraduationRequirement, 即毕业要求)"` // 毕业要求编号
	Description string `orm:"column(description);size(5000)"`                              // 指标描述
}

//-----------------------------------

// 毕业要求对培养目标支撑表
type GRSuptTrT struct {
	Id     int64 `orm:"column(id);auto"`                                             // 自增编号
	GR_id  int64 `orm:"column(GR_id);description(GR意为GraduationRequirement, 即毕业要求)"` // 毕业要求编号
	TrT_id int64 `orm:"column(TrT_id);description(TrT意为TrainTarget, 即培养目标)"`         // 培养目标编号
	Weight int   `orm:"column(weight)"`                                              // 支撑权重值
}

//-----------------------------------

// 课程信息表
type Course struct {
	Id                int64  `orm:"column(id);auto"`              // 课程自增编号
	Number            string `orm:"column(number);unique"`        // 课程号
	NameCN            string `orm:"column(name_CN)"`              // 课程中文名称
	NameEN            string `orm:"column(name_EN)"`              // 课程英文名称
	KnowledgeField    string `orm:"column(knowledge_field)"`      // 知识领域
	PreparatoryCourse string `orm:"column(preparatory_course)"`   // 先修课程
	Recommend         string `orm:"column(recommend);size(5000)"` //推荐教材和参考书目
}

// 专业包含课程对应表
type MajorMapCourse struct {
	Id        int64 `orm:"column(id);auto"`   // 自增编号
	Major_id  int64 `orm:"column(major_id)"`  // 专业编号
	Course_id int64 `orm:"column(course_id)"` // 课程编号
}

// 教师表
type Teacher struct {
	Id       int64  `orm:"column(id);auto"`       // 自增编号
	Name     string `orm:"column(name)"`          // 姓名
	Number   string `orm:"column(number);unique"` // 教师教号
	Password string `orm:"column(password)"`      // 密码
}

// 教师选课表
type TeacherSelectCourse struct {
	Id                     int64  `orm:"column(id);auto"`                   // 自增编号
	Course_id              int64  `orm:"column(course_id)"`                 // 课程编号
	Teacher_id             int64  `orm:"column(teacher_id)"`                // 教师编号
	Major_id               int64  `orm:"column(major_id)"`                  // 开课专业编号
	Term                   int    `orm:"column(term)"`                      // 开设学期
	credit                 int    `orm:"column(credit)"`                    // 学分
	TestMethod             string `orm:"column(test_method)"`               // 考核方式
	TotalPeriod            int    `orm:"column(total_period)"`              // 总学时
	TheoryPeriod           int    `orm:"column(theory_period)"`             // 理论学时
	ExperimentalPeriod     int    `orm:"column(experimental_period)"`       // 实验学时
	ComputerPeriod         int    `orm:"column(computer_period)"`           // 上机学时
	PracticePeriod         int    `orm:"column(practice_period)"`           // 实践学时
	WeekPeriod             int    `orm:"column(week_period)"`               // 周学时
	ContentRelationImgPath string `orm:"column(content_relation_img_path)"` // 教学内容关系图路径
	TeachTargetOverview    string `orm:"column(teach_target_overview)"`     // 教学目标总纲
	CourseTask             string `orm:"column(course_task)"`               // 课程性质、目的、任务
	TeachMethod            string `orm:"column(teach_method)"`              // 教学方法
	RelationOtherCourse    string `orm:"column(relation_other_course)"`     // 与其他课程的关系
	Category               string `orm:"column(category)"`                  // 课程类别(性质)
}

// 教学内容表
type TeachContent struct {
	Id                 int64  `orm:"column(id);auto"`                                             // 教学内容自增编号
	TSC_id             int64  `orm:"column(TSC_id);description(TSC意为TeacherSelectCourse, 即教师选课)"` // 教师选课编号
	Content            string `orm:"column(content);size(5000)"`                                  // 教学内容
	Demand             string `orm:"column(demand);size(5000)"`                                   // 教学要求
	MainPoint          string `orm:"column(main_point);size(5000)"`                               // 重点
	DiffcutltyPoint    string `orm:"column(diffcutly_point);size(5000)"`                          // 难点
	MethodPeriod       int    `orm:"column(method_period)"`                                       // 理论学时
	ExperimentalPeriod int    `orm:"column(experimental_period)"`                                 // 实验学时
	ComputerPeriod     int    `orm:"column(computer_period)"`                                     // 上机学时
	PracticePeriod     int    `orm:"column(practice_period)"`                                     // 实践学时
}

// 教学目标表
type TeachTarget struct {
	Id          int64  `orm:"column(id);auto"`                                             // 教学目标自增编号
	TSC_id      int64  `orm:"column(TSC_id);description(TSC意为TeacherSelectCourse, 即教师选课)"` // 教师选课编号
	Description string `orm:"column(description)"`                                         // 目标描述
}

// 教学活动表
type TeachActivity struct {
	Id     int64  `orm:"column(id);auto"`                                             // 教学活动自增编号
	TSC_id int64  `orm:"column(TSC_id);description(TSC意为TeacherSelectCourse, 即教师选课)"` // 教师选课编号
	Name   string `orm:"column(name)"`                                                // 活动名
	Weight int    `orm:"column(weight)"`                                              // 教学活动在平时成绩中占比
}

// 教学活动项表
type TeachActivityChild struct {
	Id    int64     `orm:"column(id);auto"`                                     // 教学活动项自增编号
	TA_id int64     `orm:"column(TA_id);description(TA意为TeachActivity, 即教学活动)"` // 教学活动编号
	time  time.Time `orm:"column(time)"`                                        // 活动时间
}

// 学生表
type Student struct {
	Id       int64  `orm:"column(id);auto"`       // 学生自增编号
	Number   string `orm:"column(number);unique"` // 学号
	Name     string `orm:"column(name)"`          // 姓名
	Major_id string `orm:"column(major_id)"`      // 专业编号
	Class    int    `orm:"column(class)"`         // 班级
	Grade    int    `orm:"column(grade)"`         // 年级
}

//-----------------------------------

// 教学目标对毕业指标点支撑表
type TeTSuptIP struct {
	Id     int64 `orm:"column(id);auto"`                                     // 自增编号
	TeT_id int64 `orm:"column(TeT_id);description(TeT意为TeachTarget, 即教学目标)"` // 教学目标编号
	IP_id  int64 `orm:"column(IP_id);description(IP意为IndicatorPoint, 即指标点)"` // 指标点编号
	Weight int   `orm:"column(weight)"`                                      // 权重值
}

// 教学活动项对教学目标支撑表
type TACSuptTeT struct {
	Id     int64 `orm:"column(id);auto"`                                             // 自增编号
	TAC_id int64 `orm:"column(TAC_id);description(TAC意为TeachActivityChild, 即教学活动项)"` // 教学活动项编号
	TeT_id int64 `orm:"column(TeT_id);description(TeT意为TeachTarget, 即教学目标)"`         // 教学目标编号
	Weight int   `orm:"column(weight)"`                                              // 权重值
}

// 学生选课表
type StudentSelectCourse struct {
	Id           int64 `orm:"column(id);auto"`                                             // 自增编号
	Student_id   int64 `orm:"column(student_id)"`                                          // 学生编号
	TSC_id       int64 `orm:"column(TSC_id);description(TSC意为TeacherSelectCourse, 即教师选课)"` // 教师选课编号
	ExamResult   int   `orm:"column(exam_result)"`                                         // 考试成绩
	NormalResult int   `orm:"column(normal_result)"`                                       // 平时成绩
	ExamWeight   int   `orm:"column(exam_weight)"`                                         // 考试成绩占比
	NormalWeight int   `orm:"column(normal_weight)"`                                       // 平时成绩占比
}

// 学生参与教学活动项表
type StudentJoinActivityChild struct {
	Id         int64 `orm:"column(id);auto"`                                             // 自增编号
	Student_id int64 `orm:"column(student_id)"`                                          // 学生编号
	TAC_id     int64 `orm:"column(TAC_id);description(TAC意为TeachActivityChild, 即教学活动项)"` // 教学活动项编号
	Result     int   `orm:"column(result)"`                                              // 成绩
}

//-----------------------------------

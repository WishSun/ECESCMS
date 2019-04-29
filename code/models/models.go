package models

//-----------------------------------
// 专业表
type Major struct {
	Id              string // 专业编号
	Name            string // 专业名称
	Overview        string // 专业概述
	TotalCredits    int    //总学分
	TheoryCredits   int    //理论教学学分
	PracticeCredits int    //实践教学学分
}

// 培养目标表
type TrainTarget struct {
	Id       int64  // 自增编号
	Major_id string // 专业编号
	Content  string // 目标内容
}

//-----------------------------------
// 毕业要求表
type GraduationRequirement struct {
	Id          int64  // 自增编号
	Major_id    string // 专业编号
	Name        string // 要求名称
	Description string // 要求描述
}

// 指标点表
type IndicatorPoint struct {
	Id          int64  // 自增编号
	GR_id       int64  // 毕业要求编号
	Description string // 指标描述
}

//-----------------------------------

// 毕业要求对培养目标支撑表
type GRSuptTrT struct {
	GR_id  int64 // 毕业要求编号
	TrT_id int64 // 培养目标编号
	Weight int   // 支撑权重值
}

//-----------------------------------

// 课程信息表
type Course struct {
	Id                string // 课程编号
	NameCN            string // 课程中文名称
	NameEN            string // 课程英文名称
	KnowledgeField    string // 知识领域
	PreparatoryCourse string // 先修课程
	Recommend         string //推荐教材和参考书目
}

// 专业包含课程对应表
type MajorMapCourse struct {
	Id        int64  // 自增编号
	Major_id  string // 专业编号
	Course_id string // 课程编号
}

// 教师表
type Teacher struct {
	Id       int64  // 自增编号
	Name     string // 姓名
	Account  string // 账号
	Password string // 密码
}

// 教师选课表
type TeacherSelectCourse struct {
	Course_id              string // 课程编号
	Teacher_id             int64  // 教师编号
	Major_id               string // 开课专业编号
	Term                   int    // 开设学期
	credit                 int    // 学分
	TestMethod             string // 考核方式
	TotalPeriod            int    // 总学时
	TheoryPeriod           int    // 理论学时
	ExperimentalPeriod     int    // 实验学时
	ComputerPeriod         int    // 上机学时
	PracticePeriod         int    // 实践学时
	WeekPeriod             int    // 周学时
	ContentRelationImgPath string // 教学内容关系图路径
	TeachTargetOverview    string // 教学目标总纲
	CourseTask             string // 课程性质、目的、任务
	TeachMethod            string // 教学方法
	RelationOtherCourse    string // 与其他课程的关系
	Category               string // 课程类别(性质)
}

// 教学内容表
type TeachContent struct {
	Id                 int64  // 教学内容自增编号
	Course_id          string // 课程编号
	Teacher_id         int64  // 教师编号
	Content            string // 教学内容
	Demand             string // 教学要求
	MainPoint          string // 重点
	DiffcutltyPoint    string // 难点
	MethodPeriod       int    // 理论学时
	ExperimentalPeriod int    // 实验学时
	ComputerPeriod     int    // 上机学时
	PracticePeriod     int    // 实践学时
}

// 教学目标表
type TeachTarget struct {
	Id          int64  // 教学目标自增编号
	Course_id   string // 课程编号
	Teacher_id  int64  // 教师编号
	Description string // 目标描述
}

// 教学活动表
type TeachActivity struct {
}

// 教学活动项表
type TeachActivityChild struct {
}

// 学生表
type Student struct {
}

//-----------------------------------

// 教学目标对毕业指标点支撑表
type TeTSuptIP struct {
}

// 教学活动项对教学目标支撑表
type TACSuptTeT struct {
}

// 学生选课表
type StudentSelectCourse struct {
}

// 学生参与教学活动项表
type StudentJoinActivityChild struct {
}

//-----------------------------------

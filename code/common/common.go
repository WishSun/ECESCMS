package common

type TrTType struct {
	TrTNumber  string
	TrTContent string
}

type GRIPType struct {
	GRIPNumber      string
	GRIPDescription string
}

type GRType struct {
	GRNumber      string
	GRName        string
	GRDescription string
}

type TSCType struct {
	CourseId      int64
	CourseNumber  string
	CourseName    string
	TeacherNumber string
	TeacherName   string
}

// 教师我的课程列表数据结构
type TSCListType struct {
	TSCId        int64
	CourseNumber string
	CourseName   string
	Grade        string
	MajorName    string
}

// 教学目标权值类型
type TeTWeightType struct {
	Number string
	Weight int64
}

// 给添加教学活动方法传递的教学目标权值类型
type SendTeTWeight struct {
	TeTId  string
	Weight int64
}

// 给修改教学活动方法传递教学目标权值类型
type SendTeTWeightToModify struct {
	TTWId  string
	Weight string
}

type ModifyGetTeTWeight struct {
	Number      string // 教学目标编号
	Description string // 教学目标描述
	TTWId       int64  // 教学目标在教学活动中的Id
	Weight      int64  // 教学目标在教学活动中的权值
}

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
type TSCListType struct{
	TSCId int64
	CourseNumber string
	CourseName string
	Grade string
	MajorName string
}

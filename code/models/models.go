package models

//-----------------------------------
// 专业表
type Major struct {
}

// 培养目标表
type TrainTarget struct {
}

//-----------------------------------
// 毕业要求表
type GraduationRequirement struct {
}

// 指标点表
type IndicatorPoint struct {
}

//-----------------------------------

// 毕业要求对培养目标支撑表
type GRSuptTrT struct {
}

//-----------------------------------

// 课程表
type Course struct {
}

// 教学目标表
type TeachTarget struct {
}

// 教学活动表
type TeachActivity struct {
}

// 教学活动项表
type TeachActivityChild struct {
}

// 教学内容表
type TeachContent struct {
}

// 教师表
type Teacher struct {
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

// 教师选课表
type TeacherSelectCourse struct {
}

// 学生选课表
type StudentSelectCourse struct {
}

// 学生参与教学活动项表
type StudentJoinActivityChild struct {
}

//-----------------------------------

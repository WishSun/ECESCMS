package routers

import (
	"ECESCMS/code/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/login/check_account", &controllers.LoginController{}, "post:CheckAccount")

	beego.Router("/admin_home", &controllers.AdminHomeController{})
	beego.Router("/teacher_home", &controllers.TeacherHomeController{})

	// 专业管理
	beego.Router("/admin/major_manager", &controllers.MajorManagerController{})
	// 删除专业（只接受请求url为"admin/major_manager/delete"的Get请求，并由Delete方法处理）
	beego.Router("/admin/major_manager/delete", &controllers.MajorManagerController{}, "get:Delete")
	// 基本信息管理-----------------
	beego.Router("/admin/major_base", &controllers.MajorBaseController{})
	// 基本信息添加
	beego.Router("/admin/major_base_add", &controllers.MajorBaseAddController{})
	// 基本信息修改
	beego.Router("/admin/major_base_modify", &controllers.MajorBaseModifyController{})
	// 培养目标管理-----------------
	beego.Router("/admin/major_train_target", &controllers.MajorTrainTargetController{})
	// 培养目标调整(包括修改、删除和添加)
	beego.Router("/admin/major_train_target_modify", &controllers.MajorTrainTargetModifyController{})
	// 毕业要求基本信息管理-----------------
	beego.Router("/admin/major_gr", &controllers.MajorGRController{})
	// 毕业要求指标点管理-----------------
	beego.Router("/admin/major_gr_ip", &controllers.MajorGRIPController{})
	// 毕业要求指标点调整
	beego.Router("/admin/major_gr_ip_modify", &controllers.MajorGRIPModifyController{})
	// 专业课程
	beego.Router("/admin/major_course", &controllers.MajorCourseController{})
	// 专业课程调整
	beego.Router("/admin/major_course_modify", &controllers.MajorCourseModifyController{})
	// 显示专业大纲
	beego.Router("/admin/major_view", &controllers.MajorViewController{})

	// 课程信息管理
	beego.Router("/admin/course_manager", &controllers.CourseManagerController{})
	beego.Router("/admin/course_manager/delete", &controllers.CourseManagerController{}, "get:Delete")
	beego.Router("/admin/course_add", &controllers.CourseAddController{})
	beego.Router("/admin/course_modify", &controllers.CourseModifyController{})
	beego.Router("/admin/course_view", &controllers.CourseViewController{})

	// 教师信息管理
	beego.Router("/admin/teacher_manager", &controllers.TeacherManagerController{})
	beego.Router("/admin/teacher_manager/add", &controllers.TeacherManagerController{}, "post:Add")
	beego.Router("/admin/teacher_manager/delete", &controllers.TeacherManagerController{}, "get:Delete")
	beego.Router("/admin/teacher_manager/resetpwd", &controllers.TeacherManagerController{}, "post:ResetPwd")

	// 教师选课管理
	beego.Router("/admin/teacher_select_course_manager", &controllers.TeacherSelectCourseManagerController{})
	beego.Router("/admin/teacher_select_course_view", &controllers.TeacherSelectCourseViewController{})
	beego.Router("/admin/teacher_select_course_modify", &controllers.TeacherSelectCourseModifyController{})
	// 学生信息管理
	beego.Router("/admin/student_manager", &controllers.StudentManagerController{})
	beego.Router("/admin/student_manager/add", &controllers.StudentManagerController{}, "post:Add")
	beego.Router("/admin/student_manager/delete", &controllers.StudentManagerController{}, "get:Delete")
	beego.Router("/admin/student_manager/modify", &controllers.StudentManagerController{}, "post:Modify")

	//教师我的课程管理
	beego.Router("/teacher/my_course_manager", &controllers.MyCourseManagerController{})
	beego.Router("/teacher/my_course_teach_target", &controllers.MyCourseTeachTargetController{})
	beego.Router("/teacher/my_course_teach_target/add_teach_target", &controllers.MyCourseTeachTargetController{}, "post:AddTeachTarget")
	beego.Router("/teacher/my_course_teach_target/get_gr_and_ip", &controllers.MyCourseTeachTargetController{}, "post:GetGRAndIP")
	beego.Router("/teacher/my_course_teach_target/modify_teach_target", &controllers.MyCourseTeachTargetController{}, "post:ModifyTeachTarget")
	beego.Router("/teacher/my_course_teach_target/delete", &controllers.MyCourseTeachTargetController{}, "get:DeleteTeachTarget")

	beego.Router("/teacher/my_course_modify_supt_ip", &controllers.MyCourseModifySuptIPController{})
	beego.Router("/teacher/my_course_modify_supt_ip/get_gr_ips", &controllers.MyCourseModifySuptIPController{}, "post:GetGRIPs")
	beego.Router("/teacher/my_course_modify_supt_ip/get_ip", &controllers.MyCourseModifySuptIPController{}, "post:GetIP")
	beego.Router("/teacher/my_course_base", &controllers.MyCourseBaseController{})
	beego.Router("/teacher/my_course_base_modify", &controllers.MyCourseBaseModifyController{})

	beego.Router("/teacher/my_course_activity", &controllers.MyCourseActivityController{})
	beego.Router("/teacher/my_course_activity/get_tets", &controllers.MyCourseActivityController{}, "post:GetTeTs")
	beego.Router("/teacher/my_course_activity/get_tetws", &controllers.MyCourseActivityController{}, "post:GetTeTWs")
	beego.Router("/teacher/my_course_activity/get_tet_weight", &controllers.MyCourseActivityController{}, "post:GetTeTWeight")
	beego.Router("/teacher/my_course_activity/add_activity", &controllers.MyCourseActivityController{}, "post:AddActivity")
	beego.Router("/teacher/my_course_activity/modify_activity", &controllers.MyCourseActivityController{}, "post:ModifyActivity")
	beego.Router("/teacher/my_course_activity/delete_activity", &controllers.MyCourseActivityController{}, "get:DeleteActivity")

	beego.Router("/teacher/my_course_activity_child", &controllers.MyCourseActivityChildController{})
	beego.Router("/teacher/my_course_activity_child/get_tachild_results", &controllers.MyCourseActivityChildController{}, "post:GetTAChildResults")

	beego.Router("/teacher/my_course_activity_result", &controllers.MyCourseActivityResultController{})
	beego.Router("/teacher/my_course_activity_result/get_ta_result", &controllers.MyCourseActivityResultController{}, "post:GetTAResult")

	beego.Router("/teacher/my_course_activity_child_result_manager", &controllers.MyCourseActivityChildResultManagerController{})
	beego.Router("/teacher/my_course_activity_child_result_manager/get_tachild_result", &controllers.MyCourseActivityChildResultManagerController{}, "post:GetTAChildResult")

	beego.Router("/teacher/my_course_activity_child_student_result_manager", &controllers.MyCourseActivityChildStudentResultManagerController{})
	beego.Router("/teacher/my_course_activity_child_student_result_manager/get_student_all_tachild_result", &controllers.MyCourseActivityChildStudentResultManagerController{}, "post:GetStudentAllTAChildResult")
}

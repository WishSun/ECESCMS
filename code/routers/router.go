package routers

import (
	"ECESCMS/code/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})

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

	// 学生信息管理
	beego.Router("/admin/student_manager", &controllers.StudentManagerController{})
	beego.Router("/admin/student_manager/add", &controllers.StudentManagerController{}, "post:Add")
	beego.Router("/admin/student_manager/delete", &controllers.StudentManagerController{}, "get:Delete")
	beego.Router("/admin/student_manager/modify", &controllers.StudentManagerController{}, "post:Modify")

}

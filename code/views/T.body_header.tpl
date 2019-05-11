{{define "body_header"}}
<!-- html的body体头模板 -->
<body style="background-color:black;">

<div class="navbar navbar-default top-navbar navbar-inverse" style="background-color: #222222">
    <div class="container">
        <div class="pull-left">
            <img  src="../static/img/logo.png">
        </div>
    </div>
</div>

<div class="navbar-inverse navbar-side">
    <nav class="navbar-inverse" style="width:260px;">
        <div class="sidebar-collapse">
            <ul class="nav" id="main-menu">
                <li>
                    <a href="/admin/major_manager">专业信息管理</a>
                </li>
                <li>
                    <a href="/admin/course_manager">课程信息管理</a>
                </li>
                <li>
                    <a href="/admin/teacher_manager">教师信息管理</a>
                </li>
                <li>
                    <a href="/admin/teacher_select_course_manager">教师选课管理</a>
                </li>
                <li>
                    <a href="/admin/student_manager">学生信息管理</a>
                </li>
            </ul>
        </div>
    </nav>
</div>

{{end}}

{{define "navbar"}}
<!-- 导航栏模板 -->

<a class="navbar-brand" href="/">工程认证教育体系课程管理系统</a>
<div>
      <ul class="nav navbar-nav">
         <li {{if .IsHome}}class="active"{{end}}><a href="/">关于我们</a></li>
      </ul>
</div>

<!-- 登录按钮（靠右） -->
<div class="pull-right">
    <ul class="nav navbar-nav">
        {{if .IsLogin}}  <!-- 如果已经登录, 该位置应该显示一个"退出"按钮 -->
            <li><a href="/login?quit=true">退出登录</a></li>
        {{else}} <!-- 如果没有登录, 则显示"登录"按钮 -->
            <li><a href="/login">管理员登录</a></li>
        {{end}}
    </ul>
</div>

{{end}}
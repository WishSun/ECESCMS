{{define "navbar"}}
<!-- 导航栏模板 -->

<a class="navbar-brand" href="/">我的博客</a>
<div>
      <ul class="nav navbar-nav">
         <!-- 通过在 *.go 文件中向模板传递参数，来实现不同页面使用不同的风格样式-->
         <li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
         <li {{if .IsCategory}}class="active"{{end}}><a href="/category">分类</a></li>
         <li {{if .IsTopic}}class="active"{{end}}><a href="/topic">文章</a></li>
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
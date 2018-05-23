{{define "navbar"}}
<div class="navbar navbar-default navbar-fixed-top">

	<!-- 导航 -->
	<div class="container">
		<a class="navbar-brand" href="/">我的博客</a>
		<ul class="nav navbar-nav">
			<li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
			<li {{if .IsCategory}}class="active"{{end}}><a href="/category">分类</a></li>
			<li {{if .IsTopic}}class="active"{{end}}><a href="/topic">文章</a></li>
		</ul>

		<!-- 登录退出按钮 -->
		
		<ul class="nav navbar-nav pull-right">
			{{if .IsLogin}}
			<li><a href="/login?exit=true">退出</a></li>
			{{else}}
			<li><a href="/login">登录</a></li>
			{{end}}
		</ul>
		
	</div>

</div>


{{end}}
{{define "navbar"}}

    <a class="navbar-brand" href="/">我的博客</a>
    <ur class ="nav navbar-nav">
        <li {{if .IsHome}} class = "active" {{end}}><a href="/">首页</a></li>
        <li {{if .IsCategory}} class = "active" {{end}}><a href="/category">分类</a></li>
        <li {{if .IsTop}} class = "active" {{end}}><a href="/topic">文章</a></li>
    </ur>

{{end}}
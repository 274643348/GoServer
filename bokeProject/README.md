# 项目目录结构
####  controller：控制层
####  model：数据库有关
####  views：模版文件
####  static：静态文件，存放css，js
####  conf：配置文件
####  data：嵌入式的数据库
####  log：日志信息

# 数据库
####  数据库使用ORM，而不是sql语句，推荐qbs（结构化查询）在github上

# web框架（轻量级的beego，国产）
####  支持MVC，支持REST，路由比较自由，使用简单，文档比较完善；
####  go语言其他框架比较少，国外的view比较成熟，不是轻量级，包办的东西太多；

# 前端
####  使用Bootstrap或google的开发者工具来实现


# 开发过程记录
### 第一节 创建项目目录结构和框架描述
### 第二节
#####  1.下载beego(go get github.com/astaxie/beego)
#####  2.30分钟go底层http机制
#####  3.下载bee(go get github.com/astaxie/bee)
#####  3.使用bee工具初始化Beego项目(bee new beeBokerProject)
#####  4.使用bee工具热编译Beego项目(cd beeBokerProject ;bee run;访问localhost:8080)
#####  5.conf配置的使用，日志处理的操作
### 第三节 (beego模版的使用,ORM的基本使用，view的实现)
#####  1.前20分钟简单的模版使用(beego模版：go的原生);
#####  2.分类和文章，创建数据结构，创建数据库;
1. 使用beego的ORM使用(qbs暂时不用);
2. 下载navicat premium链接数据库来浏览生成的表；
#####  3.借助bootStrap实现前端web页面的实现;
1. 下载bootstrap的css;
2. 将css相关的文件复制到static；
3. 在html中引入css(link rel="stylesheet" type="text/css" href="/static/css/bootstrap.css">)
### 第四节
#####  1.前30分钟实现登录和退出;
1. 导航栏和header独立出来；如T.header.tpl和T.navbar.tpl
2. 通过数据注入的方式来判断导航栏按钮状态（注意:html中的template中注入数据要加 .）；如：homeControl中设置IsLogin为true；
3. 通过cookie实现自动登录；操作：this.Ctx.SetCookie();（注意：重定向301是不正确的，应该用302）；
4. 登录界面实现是通过Bootstrap的css样式中的表单精简过来的;(this.Input().Get(name)获取表单数据);
5. 首页界面右方"退出"和"管理员登录"；通过检查cookie ctx.Request.Cookie(name);
6. 前端js代码控制登录界面"初步校验"和"退出"；onclick="return func();"注意return ;
#####  2.30-50实现分类的增加和删除,显示;
1. 实现分类view界面;"添加"和"删除"都是GET方式form(中增加input类型hidden，name为op，value为add )，href中直接"/category?op=del&id{{.Id}}"
2. 增加"分类"路由category.go;Get方法中获取Input().Get("op"),根据不同的类型做不同的处理
3. model模型增加数据库操作函数；CRUD:create read updata delete,参考model中的CRUM；
4. 模型定义；参考beego中的model中的模型定义；`orm:"null;auto_now;type(datetime);index"`
#####  3.go原生cookie的bug(参数中不能为空格)；

#####
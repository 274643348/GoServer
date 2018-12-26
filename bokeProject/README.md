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
#### 第一节创建项目目录结构和框架描述
#### 第二节
#####  1.下载beego(go get github.com/astaxie/beego)
#####  2.30分钟go底层http机制
#####  3.下载bee(go get github.com/astaxie/bee)
#####  3.使用bee工具初始化Beego项目(bee new beeBokerProject)
#####  4.使用bee工具热编译Beego项目(cd beeBokerProject ;bee run;访问localhost:8080)
#####  5.conf配置的使用，日志处理的操作
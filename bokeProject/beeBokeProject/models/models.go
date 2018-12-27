package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const(
	_DB_NAME = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id int64
	Title string
	Created time.Time
	Views int64
	TopicTime time.Time
	TopicCount int64
	TopicLastUserId int64
}

type Topic struct {
	Id int64
	Uid int64
	Title string
	Content string
	Attachment string
	Created time.Time
	Updated time.Time
	Views int64
	Author string
	ReplyTime time.Time
	ReplyCount int64
	RepleyLastUerId int64
}

func RegisterDB(){
	if _,err :=os.Stat(_DB_NAME); err!= nil {
		os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
		os.Create(_DB_NAME)

	}

	//注册模型
	orm.RegisterModel(new(Category),new(Topic))

	//注册驱动
	orm.RegisterDriver(_SQLITE3_DRIVER,orm.DRSqlite)

	//注册默认数据库，可以同时操作多个（必须有一个数据库default）
	orm.RegisterDataBase("default",_SQLITE3_DRIVER,_DB_NAME,10)
}

func AddCategory(name string)error{
	//获取orm
	o := orm.NewOrm()

	//创建category对象
	//cate := &Category{Title:name}
	cate := new(Category)
	cate.Title = name
	//查询判断是否已经备用
	qs :=o.QueryTable("category")
	err := qs.Filter("title",name).One(cate)
	if err == nil {
		return  err
	}

	//插入操作
	_,err =o.Insert(cate);
	if err != nil {
		return err
	}

	return nil
}

func GetAllCtegories() ([]*Category,error) {
	o := orm.NewOrm()

	cates := make([]*Category,0)

	qs := o.QueryTable("category")
	_,err := qs.All(&cates)
	return cates,err
}

func DelCategory(id string)(error){
	cid,err := strconv.ParseInt(id,10,64)
	if err != nil {
		return err
	}

	o :=orm.NewOrm()
	//删除操作和read操作都要知名主键
	cate :=&Category{Id:cid}

	_,err = o.Delete(cate);
	return err
}
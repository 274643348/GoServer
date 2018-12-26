package models

import (
	"github.com/astaxie/beego/orm"
	"os"
	"path"
	"time"
	_ "github.com/mattn/go-sqlite3"
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
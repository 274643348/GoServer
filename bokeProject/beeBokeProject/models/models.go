package models

import (
	"github.com/astaxie/beego"
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
	Title string `orm:"null"`
	Created time.Time `orm:"null;auto_now;type(datetime);index"`
	Views int64 `orm:"null"`
	TopicTime time.Time  `orm:"null;auto_now;type(date);index"`
	TopicCount int64 `orm:"null"`
	TopicLastUserId int64 `orm:"null"`
}

type Topic struct {
	Id int64
	Category string `orm:"null"`
	Uid int64 `orm:"null"`
	Title string `orm:"null"`
	Content string `orm:"null"`
	Attachment string `orm:"null"`
	Created time.Time `orm:"null;index"`
	Updated time.Time `orm:"null;index"`
	Views int64 `orm:"null"`
	Author string `orm:"null"`
	ReplyTime time.Time `orm:"null"`
	ReplyCount int64 `orm:"null"`
	RepleyLastUerId int64 `orm:"null"`
}

type Comment struct {
	Id int64
	Tid int64 `orm:"null"`
	NickName string `orm:"null"`
	Content string `orm:size(1000);null`
	Created time.Time `orm:"index;null"`
}

func RegisterDB(){
	if _,err :=os.Stat(_DB_NAME); err!= nil {
		os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
		os.Create(_DB_NAME)

	}

	//注册模型
	orm.RegisterModel(new(Category),new(Topic),new(Comment))

	//注册驱动
	orm.RegisterDriver(_SQLITE3_DRIVER,orm.DRSqlite)

	//注册默认数据库，可以同时操作多个（必须有一个数据库default）
	orm.RegisterDataBase("default",_SQLITE3_DRIVER,_DB_NAME,10)
}

//回复操作
func AddReply(tid,nickname,content string)error{
	tidNum,err := strconv.ParseInt(tid,10,64)
	if err != nil {
		return err
	}

	o :=orm.NewOrm()
	reply :=&Comment{
		Tid:tidNum,
		NickName:nickname,
		Content:content,
		Created:time.Now(),
	}
	_,err = o.Insert(reply)
	if err != nil {
		return err
	}

	//更新文章回复数
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id",tid).One(topic)
	if err == nil {
		topic.ReplyCount ++;
		topic.ReplyTime = time.Now()
		_,err = o.Update(topic)
	}

	return  err
}

func GetAllReplies(tid string)([]*Comment,error){
	tidNum,err := strconv.ParseInt(tid,10,64)
	if err != nil {
		return nil,err
	}

	o :=orm.NewOrm()
	reply :=make([]*Comment,0)
	qs :=o.QueryTable("comment")
	_,err =qs.Filter("tid",tidNum).All(&reply);
	return reply,err
}

func DeleteReply(rid string)error{
	ridNum,err := strconv.ParseInt(rid,10,64)
	if err != nil {
		return err
	}

	o :=orm.NewOrm()
	reply := &Comment{
		Id:ridNum,
	}
	if o.Read(reply) == nil {
		_,err = o.Delete(reply)
		if err != nil {
			return  err
		}
	}else {
		return err
	}

	//获取所有回复
	replies := make([]*Comment,0)
	qs := o.QueryTable("comment")
	_,err = qs.Filter("tid",reply.Tid).OrderBy("-created").All(&replies)
	if err != nil {
		return err
	}

	//更新文章回复数
	topic := new(Topic)
	qs = o.QueryTable("topic")
	err = qs.Filter("id",reply.Tid).One(topic)
	if err == nil {
		topic.ReplyCount = int64(len(replies));
		if topic.ReplyCount <=0 {
			topic.ReplyTime = time.Time{}
		}else
		{
			topic.ReplyTime = replies[0].Created;
		}
		_,err = o.Update(topic)
	}


	return  err
}
//文章操作
func AddTopic(title,category,content string)error{
	beego.Error("ljy-----------AddTopic-----title:",title,"-----content:",content)

	o := orm.NewOrm()

	topic:=&Topic{
		Title:title,
		Category:category,
		Content:content,
		Created:time.Now(),
		Updated:time.Now(),
	}

	_,err := o.Insert(topic)

	if err != nil {
		return err
	}

	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title",category).One(cate)
	if err == nil {
		cate.TopicCount ++;
		_,err = o.Update(cate)
	}

	return  err
}

func GetAllTopics(cate string ,isDesc bool)([]*Topic,error){
	beego.Error("ljy----------alltopics---cate")
	o := orm.NewOrm()

	topics := make([]*Topic,0)

	qs:=o.QueryTable("topic")

	var err error
	if isDesc {
		if len(cate) != 0 {
			qs = qs.Filter("category",cate)
		}
		_,err = qs.OrderBy("-created").All(&topics)
	}else{
		_,err = qs.All(&topics)
	}

	return topics,err
}

func GetTopic(tid string)(*Topic,error){
	tidName,err :=strconv.ParseInt(tid,10,64);
	if err != nil {
		return  nil,err
	}

	o := orm.NewOrm()

	topic := new(Topic)

	//获取topic列表
	qs := o.QueryTable("topic")

	//获取id所对应的topic
	err = qs.Filter("id",tidName).One(topic)
	if err != nil {
		return  nil,err
	}

	//增加浏览次数
	topic.Views ++;
	//更新数据
	_,err = o.Update(topic)


	return topic,err
}

func ModifyTopic(tid,title,category,content string)error{
	tidNum,err :=strconv.ParseInt(tid,10,64);
	if err != nil {
		return  err
	}

	var oldCate string
	o := orm.NewOrm()
	topic := &Topic{Id:tidNum}

	if err =o.Read(topic); err== nil{
		oldCate = topic.Category
		
		topic.Title = title
		topic.Content = content
		topic.Category = category
		topic.Updated = time.Now()
		_,err =o.Update(topic)
		if err != nil {
			return  err
		}
	}
	
	//更新分类统计
	if len(oldCate) > 0 {
		cate :=new(Category)
		qs := o.QueryTable("category")
		qs.Filter("title",oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_,err =o.Update(cate)
		}
	}

	//更新分类统计
	if len(topic.Category) > 0 {
		cate :=new(Category)
		qs := o.QueryTable("category")
		qs.Filter("title",topic.Category).One(cate)
		if err == nil {
			cate.TopicCount++
			_,err =o.Update(cate)
		}
	}




	return  nil
}
func DeleteTopic(tid string)error{
	tidNum,err :=strconv.ParseInt(tid,10,64);
	if err != nil {
		return  err
	}
	o := orm.NewOrm()
	topic :=&Topic{Id:tidNum}
	if o.Read(topic) == nil{
		_,err =o.Delete(topic)
		if err != nil {
			return err
		}
	}else{
		return err
	}

	//更新分类统计
	if len(topic.Category) > 0 {
		cate :=new(Category)
		qs := o.QueryTable("category")
		qs.Filter("title",topic.Category).One(cate)
		if err == nil {
			cate.TopicCount--
			_,err =o.Update(cate)
		}
	}
	return err

}
//分类操作
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
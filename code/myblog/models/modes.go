/**
 * 数据模型
 */
package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

/**
 * 文章分类结构
 */
type Category struct {
	Id              int64     //orm默认会将Id作为主键
	Title           string    //orm默认是255字节
	Created         time.Time `orm:"index"` //创建时间,并创建索引
	Views           int64     `orm:"index"` //浏览次数
	TopicTime       time.Time `orm:"index"` //发表时间
	TopicCount      int64     //文章数量
	TopicLastUserId int64     //最后一个操作文章的用户Id
}

/**
 * 文章
 */
type Topic struct {
	Id               int64  //文章Id
	Uid              int64  //用户Id
	Title            string //标题
	Content          string `orm:"size(5000)"`
	Category         string //文章分类
	Label            string //添加标签
	Attachment       string
	Created          time.Time `orm:"index"` //创建时间,并创建索引
	Views            int64     `orm:"index"` //浏览次数
	Updated          time.Time `orm:"index"` //更新时间
	Author           string    //作者
	ReplyTime        time.Time `orm:"index"` //回复时间
	ReplyCount       int64     //回复次数
	RepleyLastUserId int64     //最后一个回复的用户id
}

/**
 * 评论
 */
type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

/**
 * 注册数据库
 */
func RegisterDB() {

	//1.注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	//2.注册默认的数据库,注意先建立号myblog_go库
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(119.29.234.71:3306)/myblog_go?charset=utf8")

	//3.注册模型
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
}

/**
 * 添加文章分类操作
 */
func AddCategory(name string) error {
	o := orm.NewOrm()

	category := &Category{Title: name, Created: time.Now(), TopicTime: time.Now()}

	//查询文章分类名称是否已经被用了
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(category)
	if err == nil { //查不到数据，会返回错误信息
		return err
	}

	_, err = o.Insert(category)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 删除文章分类
 */
func DelCategory(id string) error {
	//10进制，64位
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	category := &Category{Id: cid}
	_, err = o.Delete(category)
	return err
}

/**
 * 获取所有文章分类
 */
func GetAllCategorys() ([]*Category, error) {
	o := orm.NewOrm()
	categorys := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&categorys)
	return categorys, err
}

/**
 * 添加文章
 */
func AddTopic(title string, category string, label string, content string, attachment string) error {
	//处理标签
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	//空格作为多个标签的分隔符
	//"beego orm"
	//[beego orm]
	//$beego#$orm#

	o := orm.NewOrm()
	topic := &Topic{
		Title:      title,
		Content:    content,
		Category:   category,
		Label:      label,
		Attachment: attachment,
		Created:    time.Now(),
		Updated:    time.Now(),
		ReplyTime:  time.Now(),
	}
	_, err := o.Insert(topic)
	if err != nil {
		return err
	}

	//更新分类
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)

	if err == nil {
		//如果不存在，简单的忽略更新操作
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	return err
}

/**
 * 获取所有文章
 */
func GetAllTopics(cate string, label string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")

	var err error
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		if len(label) > 0 {
			qs = qs.Filter("label__contains", "$"+label+"#")
		}
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

/**
 * 获取文章
 */

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)

	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)

	topic.Label = strings.Replace(strings.Replace(topic.Label, "#", " ", -1), "$", "", -1)
	return topic, err
}

/**
 * 修改文章
 */
func ModifyTopic(tid string, category string, label string, title string, content string, attachment string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	//处理标签
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	//空格作为多个标签的分隔符
	//"beego orm"
	//[beego orm]
	//$beego#$orm#

	//旧的分类 旧的附件
	var oldCate, oldAttach string

	o := orm.NewOrm()

	topic := &Topic{
		Id: tidNum,
	}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		oldAttach = topic.Attachment

		topic.Title = title
		topic.Content = content
		topic.Category = category
		topic.Label = label
		topic.Attachment = attachment
		topic.Updated = time.Now()
		_, err = o.Update(topic)
		if err != nil {
			return err
		}
	}

	//更新旧的分类统计
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}

	//删除旧的附件
	if len(oldAttach) > 0 {
		os.Remove(path.Join("attachment", oldAttach))
	}

	//更新新的分类统计
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	}

	return nil
}

/**
 * 删除文章
 */
func DeleteTopic(id string) error {
	//10进制，64位
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	//旧的分类
	var oldCate string

	o := orm.NewOrm()
	topic := &Topic{Id: tid}

	if o.Read(topic) == nil {
		oldCate = topic.Category
		_, err = o.Delete(topic)

		if err != nil {
			return err
		}
	}

	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}

	return err
}

/**
 * 添加评论
 */
func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)

	if err != nil {
		return err
	}

	reply := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)

	if err != nil {
		return err
	}

	//更新文章的评论数和评论时间
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = time.Now()
		topic.ReplyCount++
		_, err = o.Update(topic)
	}
	return err
}

/**
 * 获取所有的评论
 */
func GetAllReplies(tid string) (replies []*Comment, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	replies = make([]*Comment, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_, err = qs.Filter("Tid", tidNum).All(&replies)
	if err != nil {
		return nil, err
	}
	return replies, err
}

/**
 * 删除评论
 */
func DeleteReply(rid string) error {
	//10进制，64位
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	var tidNum int64

	comment := &Comment{Id: ridNum}

	if o.Read(comment) == nil {
		tidNum = comment.Tid
		_, err = o.Delete(comment)
		if err != nil {
			return err
		}
	}

	replies := make([]*Comment, 0)
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)
	if err != nil {
		return err
	}

	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		if len(replies) != 0 {
			topic.ReplyTime = replies[0].Created
			topic.ReplyCount = int64(len(replies))
			_, err = o.Update(topic)
		} else {
			topic.ReplyCount = 0
			_, err = o.Update(topic)
		}
	}
	return err
}

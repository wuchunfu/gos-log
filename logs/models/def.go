package models

import (
	"fmt"
	"time"

	log "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//TClient 客户端
type TClient struct {
	Id          int64     `json:"id" pk:"auto" orm:"column(id)"`
	Ip          string    `json:"ip" orm:"column(ip)"`
	Port        string    `json:"port" orm:"column(port)"`
	Vkey        string    `json:"vkey" orm:"column(vkey)"`
	Info        string    `json:"info" orm:"column(info)"`
	Zip         string    `json:"zip" orm:"column(zip)"`
	Online      string    `json:"online" orm:"column(online)"`
	Status      string    `json:"status" orm:"column(status)"`
	CreatedBy   string    `json:"createdBy" orm:"column(created_by)"`
	CreatedTime time.Time `json:"createdTime" orm:"column(created_time)"`
	UpdatedBy   string    `json:"updatedBy" orm:"column(updated_by)"`
	UpdatedTime time.Time `json:"updatedTime" orm:"column(updated_time)"`
}

//TItem 客户端
type TItem struct {
	Id          int64     `json:"id" pk:"auto" orm:"column(id)"`
	ClientId    int64     `json:"clientId" orm:"column(client_id)"`
	ItemName    string    `json:"itemName" orm:"column(item_name)"`
	ItemDesc    string    `json:"itemDesc" orm:"column(item_desc)"`
	LogPath     string    `json:"logPath" orm:"column(log_path)"`
	LogPrefix   string    `json:"logPrefix" orm:"column(log_prefix)"`
	LogSuffix   string    `json:"logSuffix" orm:"column(log_suffix)"`
	Status      string    `json:"status" orm:"column(status)"`
	CreatedBy   string    `json:"createdBy" orm:"column(created_by)"`
	CreatedTime time.Time `json:"createdTime" orm:"column(created_time)"`
	UpdatedBy   string    `json:"updatedBy" orm:"column(updated_by)"`
	UpdatedTime time.Time `json:"updatedTime" orm:"column(updated_time)"`
}

//Page 分页
type Page struct {
	PageNo     int         `json:"pageNo"`
	PageSize   int         `json:"pageSize"`
	TotalPage  int         `json:"totalPage"`
	TotalCount int         `json:"totalCount"`
	FirstPage  bool        `json:"firstPage"`
	LastPage   bool        `json:"lastPage"`
	List       interface{} `json:"list"`
}

//DBConfig 数据相关配置
type DBConfig struct {
	Host         string
	Port         string
	Database     string
	Username     string
	Password     string
	MaxIdleConns int //最大空闲连接
	MaxOpenConns int //最大连接数
}

type Def struct {
	DBConf *DBConfig
}

/*
 * clientmanger构造器
 */
func NewDef(dbConf *DBConfig) *Def {
	mgr := &Def{
		DBConf: dbConf,
	}
	//初始化orm
	mgr.initDB()
	return mgr
}

/**
  初始化db，注册默认数据库，同时将实体模型也注册上去
*/
func (mgr *Def) initDB() {
	// 是否开启调试模式 调试模式下会打印出sql语句
	orm.Debug = true
	// orm.RegisterDriver("postgres", orm.DRPostgres)
	ds := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", mgr.DBConf.Host, mgr.DBConf.Port, mgr.DBConf.Username, mgr.DBConf.Password, mgr.DBConf.Database)
	log.Info("datasource=[%s]", ds)
	// err := orm.RegisterDataBase("default", "postgres", ds, mgr.DBConf.MaxIdleConns, mgr.DBConf.MaxOpenConns)
	// err := orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/logs?charset=utf8&parseTime=true&loc=Local")
	err := orm.RegisterDataBase("default", "mysql", mgr.DBConf.Username+":"+mgr.DBConf.Password+"@tcp("+mgr.DBConf.Host+":"+mgr.DBConf.Port+")/"+mgr.DBConf.Database+"?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	orm.RegisterModel(new(TClient), new(TItem))
}

//PageUtil 分页工具
func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{
		PageNo:     pageNo,
		PageSize:   pageSize,
		TotalPage:  tp,
		TotalCount: count,
		FirstPage:  pageNo == 1,
		LastPage:   pageNo == tp,
		List:       list,
	}
}

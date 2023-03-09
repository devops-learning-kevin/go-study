package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"time"
)

var DB *gorm.DB //后续用于操作数据库

// User 定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
type User struct {
	ID int64 // 主键
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	//例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
	CreateTime time.Time `gorm:"column:createtime"`
}

type Address struct {
	ID         int64
	Addr       string    `gorm:"column:address"`
	Balance    string    `gorm:"column:balance"`
	Remark     string    `gorm:"remark"`
	CreateTime time.Time `gorm:"column:createtime"`
}

func (u *User) TableName() string {
	return "users" //你希望他的名字是什么就填什么
}

func (a *Address) TableName() string {
	return "address"
}

func GetForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var addr string
	var bal string
	for k, v := range r.Form {
		fmt.Printf("key is %T,%v\n: ", k, k)
		fmt.Printf("val is:%T,%v\n ", v, v)
		if k == "address" {
			addr = v[0]
		}
		if k == "balance" {
			bal = v[0]
		}

	}

	if !(DB.HasTable(&Address{})) {

		DB.AutoMigrate(&Address{})
	}
	if addr != "" {
		address := Address{Addr: addr, Balance: bal, Remark: "remak1", CreateTime: time.Now()}
		//DB.NewRecord(address)
		DB.Create(&address)

		fmt.Fprintf(w, "保存成功！")
	}
}

func init() {
	//配置MySQL连接参数
	username := "root"       //账号
	password := "Ck834121!"  //密码
	host := "110.41.154.234" //数据库地址，可以是Ip或者域名
	port := 3306             //数据库端口
	Dbname := "gorm"         //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	DB = db
	fmt.Println("连接数据库成功！")
}

func main() {
	fmt.Println("Server started ... Listening")
	http.HandleFunc("/", GetForm)
	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

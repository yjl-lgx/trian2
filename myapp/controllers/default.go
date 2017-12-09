package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"net"
	"crypto/md5"
    "encoding/hex"
	"os"
	"database/sql"
	"time"
    _"github.com/go-sql-driver/mysql" 
)
//全局
var db *sql.DB
var uname string 
var intoduce string 
var myhe string
var myha string
//错误
func checkErr(err error) {
	if err != nil {
       panic(err)
	}
}
//log信息 
func myhome(){
	fmt.Println("xxxxxxxxxxxxxx")
	myhe = os.Getenv("OEM")
	myha = os.Getenv("VER")
	fmt.Println(myhe)
	fmt.Println(myha)
}
//首页
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//登录
type Zhmyapp struct {
	beego.Controller
}

func (c *Zhmyapp) LogA() {
	c.TplName = "zhmyapp.tpl"
	
}

//获取登录数据
func (c *Zhmyapp) LogB() {
	c.TplName = "zhmyapp.tpl"
	myhome()
	username:= c.GetString("username")
    if username == "" {
        c.Ctx.WriteString("username is empty")
        return
	}
	password:= c.GetString("password")
    if password == "" {
        c.Ctx.WriteString("password is empty")
        return
	}
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.34.20:3306)/dbname?charset=utf8")
	if err != nil {
        fmt.Println("打开SQL时出错:", err.Error())
        return
	}
	//密码转换
	h := md5.New()
	h.Write([]byte(password))
	ph := hex.EncodeToString(h.Sum(nil))
	//匹配帐号密码
    var id    int 
	var pword string
	var te   string
	rows, err := db.Query("select * from runoob_tbl where BINARY runoob_password=? and runoob_username=?",ph,username)
	defer rows.Close()
	checkErr(err)
	for rows.Next() {
		//账号和密码匹配
		rows.Columns()
		err = rows.Scan(&id, &uname, &pword, &intoduce,&te)
		checkErr(err)	
		c.Redirect("/user/profile",302)
	    //fmt.Println("信息",id,uname,intoduce,pword,te)
		//myhome()
	}	
	defer db.Close()		
	//帐号密码不匹配
	c.Redirect("/user/signup",302)	
}

//注册
type Sigmyapp struct {
	beego.Controller
}

func (c *Sigmyapp) SigA() {
	
	c.TplName = "Sigmyapp.tpl"
}

//填写注册信息
func (c *Sigmyapp) SigB() {
	//获取网页信息
	c.TplName = "Sigmyapp.tpl"
//	fmt.Println("haha")
	Username:= c.GetString("Username")
	fmt.Println("username",Username)
    if Username == "" {
        c.Ctx.WriteString("Username is empty")
        return
	}
	Password:= c.GetString("password")
	fmt.Println("password",Password)
    if Password == "" {
        c.Ctx.WriteString("password is empty")
        return
	}
	Introdname:= c.GetString("introdname")
    if Introdname == "" {
        c.Ctx.WriteString("introdname is empty")
        return 
	}
//	fmt.Println("introdname",Introdname)
	//转换密码
	h := md5.New()
	h.Write([]byte(Username))
//	fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil)))
	//当前时间 
	t := time.Now()
	now := t.Format("2006-01-02 03:04:05 PM")
    fmt.Println(now)
	//数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.34.20:3306)/dbname?charset=utf8")
	if err != nil {
        fmt.Println("打开SQL时出错:", err.Error())
        return
	}
	//创建数据库中的数据表
/*
CREATE TABLE IF NOT EXISTS `runoob_tbl`(
   `runoob_id` INT UNSIGNED AUTO_INCREMENT,
   `runoob_username` VARCHAR(40) NOT NULL,
   `runoob_password` VARCHAR(100) NOT NULL,
   `runoob_introdname` VARCHAR(100) NOT NULL,
   `submission_date` VARCHAR(100) NOT NULL,
   PRIMARY KEY ( `runoob_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
	//检查数据是否存在
	rows, err := db.Query("select runoob_username from runoob_tbl where BINARY runoob_username=?",Username)
	//checkErr(err) 
	//数据出错
	fmt.Println(rows)
    if rows.Next() == false {
		//如果数据库中没有数据将数据插入到表中
		stmt, err := db.Prepare("INSERT runoob_tbl SET runoob_username=?,runoob_password=?,runoob_introdname=?,submission_date=?")  
		checkErr(err)  
		_, err = stmt.Exec(Username, hex.EncodeToString(h.Sum(nil)), Introdname,now)  
		checkErr(err) 
		defer db.Close()
		//结束后跳转到登陆页面
		c.Redirect("/user/login",302)
    }else{
		//数据库中有数据则报错
        c.Ctx.WriteString("The user already exists")
        return 
	}
	defer rows.Close()
}

//app页
type Myapp struct {
	beego.Controller
}
func (m *Myapp) Get() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var ipnet *net.IPNet
	ipnet = addrs[1].(*net.IPNet)
	m.Data["User"] = m.Ctx.Input.UserAgent()
	m.Data["My_app"] = ipnet
	m.Data["Products"] = myhe
	m.Data["Edition"] = myha
	m.Data["username"] = uname
	m.Data["introduce"] = intoduce
	m.Data["My_hostname"],_ = os.Hostname()
	m.Data["Website"] = "beego.me"
	m.Data["Email"] = "liguoxi@ghostcloud.cn"
	m.TplName = "myapp.tpl"
}




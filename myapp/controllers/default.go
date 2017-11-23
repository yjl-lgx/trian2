package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"net"
	"os"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

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
	//var ok bool
	/*for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok = address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("ipv4:",ipnet.IP.String())
			}
	
	}*/
	ipnet = addrs[1].(*net.IPNet)
	m.Data["User"] = m.Ctx.Input.UserAgent()
	m.Data["My_app"] = ipnet
	m.Data["My_hostname"],_ = os.Hostname()
	m.Data["Website"] = "beego.me"
	m.Data["Email"] = "liguoxi@ghostcloud.cn"
	m.TplName = "myapp.tpl"
}

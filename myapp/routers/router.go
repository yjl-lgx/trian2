package routers

import (
	"myapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/profile", &controllers.Myapp{})
	beego.Router("/user/login", &controllers.Zhmyapp{},"get:LogA;post:LogB")
	beego.Router("/user/signup", &controllers.Sigmyapp{},"get:SigA;post:SigB")
}
 
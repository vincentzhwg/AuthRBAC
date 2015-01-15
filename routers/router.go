package routers

import (
	"github.com/vincentzhwg/authrbac/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

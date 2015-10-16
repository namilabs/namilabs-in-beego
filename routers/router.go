package routers

import (
	"namilabs/controllers"
	"github.com/astaxie/beego"
)

func init() {
  beego.Router("/", &controllers.MainController{})
  beego.Router("/home", &controllers.MainController{})
  beego.Router("/nladmin", &controllers.AdminController{}, "get:Dashboard")
  
  
  //auth
  beego.Router("/user/login/:back", &controllers.UserController{}, "get,post:Login")
  beego.Router("/auth/login", &controllers.UserController{}, "get,post:Login")
  beego.Router("/auth/logout", &controllers.UserController{}, "get:Logout")

  //user
  beego.Router("/user/register", &controllers.UserController{}, "get,post:Register")
  beego.Router("/user/profile", &controllers.UserController{}, "get,post:Profile")
  beego.Router("/user/verify/:uuid([0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12})", &controllers.UserController{}, "get:Verify")

  //category
  beego.Router("/nladmin/category", &controllers.CategoryController{}, "get:Index")
  beego.Router("/nladmin/category/add", &controllers.CategoryController{}, "get:GetAdd")
  beego.Router("/nladmin/category/add", &controllers.CategoryController{}, "post:PostAdd")
}

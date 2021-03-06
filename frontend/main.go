package main

import (
	"github.com/astaxie/beego"
	"github.com/wwgberlin/bubble/frontend/controllers"
	"os"
)

func main() {
	os.MkdirAll("./images/", os.ModePerm)
	beego.Router("/", &controllers.UsersController{})
	beego.SetStaticPath("/images", "images")
	beego.Run()
}

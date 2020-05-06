package main

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (controller *MainController) Get() {
	controller.Ctx.WriteString("感谢欧阳老板给我机会，尽管老板你姓欧")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}

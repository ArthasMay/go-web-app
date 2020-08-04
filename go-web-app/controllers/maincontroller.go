package controllers

import (
	"beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Username"] = "ArthasMay"
	this.Data["Email"] = "ArthasMay@gmail.com"
	this.TplNames = "index.tpl"
}

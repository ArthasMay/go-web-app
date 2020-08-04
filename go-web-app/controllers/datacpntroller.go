package controllers

import (
	"beego"
)

type DataController struct {
	beego.Controller
}

type LIKE struct {
	Food   string
	Watch  string
	Listen string
}

type JSONS struct {
	Code string
	Msg  string
	User []string `json:"user_info"`
	Like LIKE
}

func (c *DataController) Get() {
	data := &JSONS{"100", "获取成功",
		[]string{"maple", "18"}, LIKE{"蛋糕", "电影", "音乐"}}
	c.Data["json"] = data
	c.ServeJSON()
}

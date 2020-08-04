package beego

import (
	"beego/context"
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type Controller struct {
	Ctx       *context.Context
	Tpl       *template.Template
	Data      map[interface{}]interface{}
	ChildName string
	Layout    []string
	TplExt    string
	TplNames  string

	ViewPath string
}

type ControllerIntrerface interface {
	Init(ctx *context.Context, cn string)
	Prepare()
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
	Finish()
	Render() error
}

func (c *Controller) Init(ctx *context.Context, cn string) {
	c.Data = make(map[interface{}]interface{})
	c.Layout = make([]string, 0)
	c.TplNames = ""
	c.Ctx = ctx
	c.ChildName = cn
	c.TplExt = "tpl"
	c.ViewPath = "./beego"
}

func (c *Controller) Prepare() {

}

func (c *Controller) Finish() {

}

func (c *Controller) Get() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Post() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Delete() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Put() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Head() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Patch() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Options() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Render() error {
	if len(c.Layout) > 0 {
		var filenames []string
		for _, file := range c.Layout {
			filenames = append(filenames, path.Join(c.ViewPath, file))
		}
		t, err := template.ParseFiles(filenames...)
		if err != nil {
			fmt.Printf("template parsefiles err: %v", err)
		}
		err = t.ExecuteTemplate(c.Ctx.ResponseWriter, c.TplNames, c.Data)
		if err != nil {
			fmt.Printf("template execute err:%v", err)
		}
	} else {
		if c.TplNames == "" {
			c.TplNames = c.ChildName + "/" + c.Ctx.Request.Method + c.TplExt
		}

		t, err := template.ParseFiles(path.Join(c.ViewPath, c.TplNames))
		if err != nil {
			fmt.Printf("template parsefiles err: %v", err)
		}
		err = t.Execute(c.Ctx.ResponseWriter, c.Data)
		if err != nil {
			fmt.Printf("template execute err:%v", err)
		}
	}
	return nil
}

func (c *Controller) Redirect(url string, code int) {

}

func (c *Controller) ServeJSON(encoding ...bool) {
	var (
		hasIndent = false
		hasEncoding = false
	)

	c.Ctx.Output.JSON(c.Data["json"], hasIndent, hasEncoding)
}

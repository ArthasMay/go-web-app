package beego

import (
	"beego/context"
)

type Controller struct {
	Ctx       *context.Context
	ChildName string
}

type ControllerIntrerface interface {
	Init(ctx context.Context, cn string)
	Prepare()
	Get()
	Post()
}
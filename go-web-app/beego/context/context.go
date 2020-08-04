package context

import (
	"time"
	"net/http"
)

type Response struct {
	http.ResponseWriter
	Started bool
	Status int
	Elapsed time.Duration
}


type Context struct {
	Output		   *BeegoOutput
	Request		   *http.Request
	ResponseWriter *Response
}

func (ctx *Context) Reset(rw http.ResponseWriter, r *http.Request) {
	ctx.Request = r
	if ctx.ResponseWriter == nil {
		ctx.ResponseWriter = &Response{}
	}
	ctx.ResponseWriter.reset(rw)
	ctx.Output.Reset(ctx)
}

func (r *Response) reset(rw http.ResponseWriter) {
	r.ResponseWriter = rw
	r.Status = 0
	r.Started = false
}
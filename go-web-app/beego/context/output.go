package context

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type BeegoOutput struct {
	Context    *Context
	Status     int
	EnableGzip bool
}

func NewOutput() *BeegoOutput {
	return &BeegoOutput{}
}

func (output *BeegoOutput) Reset(ctx *Context) {
	output.Context = ctx
	output.Status = 0
}

func (output *BeegoOutput) Header(key, val string) {
	output.Context.ResponseWriter.Header().Set(key, val)
}

func (output *BeegoOutput) Body(content []byte) error {
	var encoding string
	var buf = &bytes.Buffer{}
	if output.EnableGzip {
		encoding = ParseEncoding(output.Context.Request)
	}
	if b, n, _ := WriteBody(encoding, buf, content); b {
		output.Header("Content-Encoding", n)
		output.Header("Content-Length", strconv.Itoa(buf.Len()))
	} else {
		output.Header("Content-Length", strconv.Itoa(len(content)))
	}
	// Write status code if it has been set manually
	// Set it to 0 afterwards to prevent "multiple response.WriteHeader calls"
	if output.Status != 0 {
		output.Context.ResponseWriter.WriteHeader(output.Status)
		output.Status = 0
	} else {
		output.Context.ResponseWriter.Started = true
	}
	io.Copy(output.Context.ResponseWriter, buf)
	return nil
}

func (output *BeegoOutput) JSON(data interface{}, hasIndent bool, encoding bool) error {
	output.Header("Content-Type", "application/json; charset=utf-8")
	var content []byte
	var err error
	content, err = json.Marshal(data)

	if err != nil {
		http.Error(output.Context.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}

	return output.Body(content)
}

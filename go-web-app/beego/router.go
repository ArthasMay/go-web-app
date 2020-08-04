package beego

import (
	"beego/context"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
)

type controllerInfo struct {
	regex          *regexp.Regexp
	params         map[int]string
	controllerType reflect.Type
}

type ControllerRegister struct {
	routers []*controllerInfo
}

func NewControllerRegister() *ControllerRegister {
	return &ControllerRegister{
		routers: []*controllerInfo{},
	}
}

func (p *ControllerRegister) Add(pattern string, c ControllerIntrerface) {
	parts := strings.Split(pattern, "/")

	j := 0
	params := make(map[int]string)
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			expr := "([^/]+)"

			//a user may choose to override the defult expression
			// similar to expressjs: ‘/user/:id([0-9]+)’
			if index := strings.Index(part, "("); index != -1 {
				expr = part[index:]
				part = part[:index]
			}
			params[j] = part
			parts[i] = expr
			j++
		}

		pattern = strings.Join(parts, "/")
		regex, regexErr := regexp.Compile(pattern)
		if regexErr != nil {
			// TODO: add error handling here to avoid panic
			panic(regexErr)
		}

		t := reflect.Indirect(reflect.ValueOf(c)).Type()
		route := &controllerInfo{}
		route.regex = regex
		route.params = params
		route.controllerType = t

		p.routers = append(p.routers, route)
	}
}

var RecoverPanic = false

func (p *ControllerRegister) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var started bool
	for prefix, staticDir := range BConfig.WebConfig.StaticDir {
		if strings.HasPrefix(r.URL.Path, prefix) {
			file := staticDir + r.URL.Path[len(prefix):]
			http.ServeFile(w, r, file)
			started = true
			return
		}
	}

	requestPath := r.URL.Path

	// find a matching route
	for _, route := range p.routers {

		// check if route pattern matches url
		if !route.regex.MatchString(requestPath) {
			continue
		}

		// get submatches()
		matches := route.regex.FindStringSubmatch(requestPath)

		// double check that the route matches the url pattern
		if len(matches[0]) != len(requestPath) {
			continue
		}

		params := make(map[string]string)
		if len(route.params) > 0 {
			// add url parameters to the query param map
			values := r.URL.Query()

			for i, match := range matches[1:] {
				values.Add(route.params[i], match)
				params[route.params[i]] = match
			}
			// reassemble query params and add to RawQuery
			r.URL.RawQuery = url.Values(values).Encode() + "&" + r.URL.RawQuery
		}

		// Invoke the request handler
		vc := reflect.New(route.controllerType)
		init := vc.MethodByName("Init")
		in := make([]reflect.Value, 2)
		ct := &context.Context{ResponseWriter: &context.Response{ResponseWriter: w, Started: true, Status: 200, Elapsed: time.Duration(10)}, Request: r, Output: context.NewOutput()}
		ct.Reset(w, r)
		in[0] = reflect.ValueOf(ct)
		in[1] = reflect.ValueOf(route.controllerType.Name())
		init.Call(in)
		in = make([]reflect.Value, 0)
		method := vc.MethodByName("Prepare")
		method.Call(in)
		if r.Method == "GET" {
			method = vc.MethodByName("Get")
			method.Call(in)
		} else if r.Method == "POST" {
			method = vc.MethodByName("Post")
			method.Call(in)
		} else if r.Method == "HEAD" {
			method = vc.MethodByName("Head")
			method.Call(in)
		} else if r.Method == "DELETE" {
			method = vc.MethodByName("Delete")
			method.Call(in)
		} else if r.Method == "PUT" {
			method = vc.MethodByName("Put")
			method.Call(in)
		} else if r.Method == "PATCH" {
			method = vc.MethodByName("Patch")
			method.Call(in)
		} else if r.Method == "OPTIONS" {
			method = vc.MethodByName("Options")
			method.Call(in)
		}

		if !ct.ResponseWriter.Started && ct.ResponseWriter.Status == 0 {
			if BConfig.WebConfig.AutoRender {
				method = vc.MethodByName("Render")
				method.Call(in)
			}
		}

		method = vc.MethodByName("Finish")
		method.Call(in)
		started = true
		break
	}

	if started == false {
		http.NotFound(w, r)
	}
}

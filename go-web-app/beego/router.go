package beego

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

type controllerInfo struct {
	regex          *regexp.Regexp
	params         map[int]string
	controllerType reflect.Type
}

type ControllerRegistor struct {
	routers     []*controllerInfo
	Application *App
}

func (p *ControllerRegistor) Add(pattern string, c ControllerIntrerface) {
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

func (p *ControllerRegistor) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		if !RecoverPanic {
	// 			panic(err)
	// 		} else {
	// 			Critical()
	// 		}
	// 	}
	// }()

	var started bool
	fmt.Println(started)
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
		

	}

}

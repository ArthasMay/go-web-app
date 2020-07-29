package main

import (
	"fmt"
	"regexp"
)

func main() {
	regexDemo()
}

func demoDefer() {
	defer func() {
		fmt.Println("defer exec")
	}()

	fmt.Println("Hello")
}

func regexDemo() {
	pattern := "/user/:([^/]+)/([0-9])"

	regex, _ := regexp.Compile(pattern)

	if regex.MatchString("/user/:100/name") {
		fmt.Println("success")
	}

	matches := regex.FindStringSubmatch("/user/:100/123")
	fmt.Println(matches)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "aaaaaa"
	var a = "hello, "
	var b = "world"
	var c = a + b
	fmt.Println(len(str))
	fmt.Println(c)
	var d = strings.Split(c, ",")
	fmt.Println(d)
}

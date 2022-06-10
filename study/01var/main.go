package main

import "fmt"

// 声明变量，驼峰式命名
var studentName string

var (
	name string
	age  int
	isOk bool
)

func main() {
	studentName = "zhangsan"
	name = "lisi"
	age = 10
	isOk = false
	var a = 18
	fmt.Println(studentName)
	fmt.Println(name)
	fmt.Printf("年龄是%d\n", age)
	fmt.Println(isOk)
	fmt.Println(a)
	fmt.Println("Hello World")
}

package main

import "fmt"

const a = 100

const (
	b = 10
	c = 20
	d = "a"
	e
	f
	n1 = iota
)

const n2 = iota
const n3 = iota
const (
	n4 = iota
	n5
	n6
)

// 每行常量声明会让iota增加1
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

func main() {
	fmt.Println(a)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
}

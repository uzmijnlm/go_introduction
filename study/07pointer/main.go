package main

import "fmt"

func main() {

	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Printf("%T\n", p)

	m := *p
	fmt.Println(m)

	fmt.Println(&*p)
	fmt.Println(*&*p)

	var a = new(int)
	*a = 100
	fmt.Println(*a)

}

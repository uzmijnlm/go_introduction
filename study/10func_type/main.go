package main

import "fmt"

func f1() {
	fmt.Println("f1")
}

func f2() int {
	fmt.Println("f2")
	return 10
}

func f3(x func()) {
	fmt.Println("f3")
	x()
}

func f4(x func() int) {
	fmt.Println("f6")
	ret := x()
	fmt.Println(ret)
}

func f5(x func() int) func() {
	fmt.Println(x())
	return func() {
		fmt.Println("inner func")
	}
}

func bibao(m, n int) func(x func(int, int) int) int {
	return func(x func(int, int) int) int {
		return x(m, n)
	}
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)

	b := f2
	fmt.Printf("%T\n", b)

	f3(f1)

	f4(f2)

	f5(f2)()

	f6 := func() int {
		return 20
	}

	f5(f6)()

	var xxx = bibao(3, 4)
	var ret = xxx(func(a int, b int) int {
		i := a * b
		return i
	})
	fmt.Println(ret)

}

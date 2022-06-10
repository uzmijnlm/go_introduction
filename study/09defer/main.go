package main

import "fmt"

// return不是原子操作，会先执行返回值赋值，再return返回值，如果有defer的话会在两者之间执行
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值定义出没有定义返回值的变量名，所以这里是把x的值赋值给返回值，那么返回值就是5。defer修改了x的值，但是不影响返回值5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值定义为了x，所以这一步会把5赋值给x，然后执行defer，x增加了1，最后返回的还是x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	r1 := f1()
	fmt.Println(r1)

	r2 := f2()
	fmt.Println(r2)

	r3 := f3()
	fmt.Println(r3)

	r4 := f4()
	fmt.Println(r4)

}

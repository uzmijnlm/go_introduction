package main

import "fmt"

func sum(x int, y int) (ret int) {
	return x + y
}

func sum2(x int, y int) (ret int) {
	ret = x + y
	return
}

func f2() {
	fmt.Println("f2")
}

// 返回值
func f3() int {
	return 10
}

func f5() (int, int) {
	return 1, 2
}

func f6(x, y int) int {
	return x - y
}

func f7(x, y int, z ...string) int {
	return x - y
}

func main() {

	r1 := sum(10, 20)
	fmt.Println(r1)

}

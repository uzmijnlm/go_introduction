package main

import "fmt"

func main() {
	// 必须指定存放的元素的类型和容量
	// 数组的长度是不能变的，是类型的一部分，[3]int与[2]int是不一样的类型
	var a [3]bool

	a = [3]bool{true, true, false}

	fmt.Println(a)

	a1 := [...]int{0, 1, 2}
	fmt.Println(a1)

	a2 := [5]int{3, 3: 5}
	fmt.Println(a2)
	fmt.Println(a2[3])

	var a3 [3][2]string
	a3 = [3][2]string{
		{"a", "b"},
		[2]string{"c"},
	}
	fmt.Println(a3)

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1)
	fmt.Println(b2)

}

package main

import "fmt"

func main() {

	var s []int
	s = []int{1, 2, 3}
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = append(s, 4)
	fmt.Println(cap(s))
	s = append(s, 5)
	fmt.Println(cap(s))
	s = append(s, 6)
	fmt.Println(cap(s))
	s = append(s, 7)
	fmt.Println(cap(s))

	a1 := [...]int{1, 2, 3, 4, 5}
	s1 := a1[0:4]
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	s2 := a1[3:4]
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	s3 := a1[3:5]
	fmt.Println(cap(s3))
	s2 = append(s2, 0)
	s2 = append(s2, 0)
	s2 = append(s2, 0)
	s2 = append(s2, 0)
	fmt.Println(cap(s2))
	fmt.Println(s2)

	s4 := s2[3:4]
	fmt.Println(cap(s4))

	a := [...]int{1, 2, 3, 4, 5}
	b := a[0:1]
	b[0] = 2
	fmt.Println(a)

	// 使用make函数创建切片
	slice1 := make([]int, 5, 10)
	fmt.Println(slice1)

	var slice2 = make([]int, 3, 3)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	slice2[0] = 1
	fmt.Println(slice1, slice2)

}

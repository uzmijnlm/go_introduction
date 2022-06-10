package main

import "fmt"

func funcA() {
	fmt.Println("a")
}

func funcB() {
	defer func() {
		err := recover() // recover必须搭配defer使用，defer语句一定要在可能引发panic的语句之前定义
		fmt.Println(err)
		fmt.Println("释放连接")
	}()
	panic("严重错误") // 程序崩溃退出
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()

}

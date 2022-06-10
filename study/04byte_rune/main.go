package main

import "fmt"

func main() {
	s := "aaaabbbb"

	s1 := []rune(s)
	fmt.Println(s1)
	s1[0] = 'b'
	fmt.Println(string(s1))

	s2 := "big"
	s3 := []byte(s2)
	s3[0] = 'c'
	fmt.Println(string(s3))
}

package main

import "fmt"

const (
	BEIJING = iota // iota每次累加1，相当于enum用法
	SHANGHAI
	HANGZHOU
	NANJING
	CHENGDU
)

func main() {
	const length int = 10
	fmt.Println("length =", length)

	// 常量只可读
	//length = 100

	/*iota只能配合const() 进行累加
	var a int = iota
	println("a =", a)*/

}

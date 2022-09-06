package main

import "fmt"

func printType(arg interface{} /* 空接口，可以包含任何对象，相当于任何对象都实现了该接口 */) {
	fmt.Println("printType()...")
	fmt.Println("arg is", arg)

	// 类型断言，只能配合接口使用
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg is string type")
		fmt.Println("arg's value is ", value)
	} else {
		fmt.Println("arg is not string type")
	}
	fmt.Println("==========================")
}

type Book struct {
	title string
}

func main() {
	var book Book
	printType(book)

	printType(100)
	printType("Hello Golang")
	printType(true)
}

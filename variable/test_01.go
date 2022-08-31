package main

import (
	"fmt"
)

// 使用var是可以声明全局变量
var gA int = 100
var gB = 200

// :=只能使用在函数体内
//gC := 300 // 'gC' unexpected

func main() {
	// 声明一个变量默认值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("a = %T\n", a)

	// 声明一个变量并赋予初始值为200
	var b int = 200
	fmt.Println("b = ", b)
	fmt.Printf("b = %T\n", b)

	// 声明一个变量指定数据类型为字符串
	var c string = "Hello GoLang!"
	fmt.Println("c = ", c)
	fmt.Printf("c = %T\n", c)
	fmt.Printf("c = %s\n", c)

	// 声明一个变量自动推导其数据类型
	var d = true
	fmt.Println(d)

	// 声明及其初始化一个变量
	e := 200
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T", e)

	// 声明多个变量
	var aa, bb int = 100, 200
	fmt.Println("aa = ", aa, "bb = ", bb)

	var cc, dd = 300, "Hello GoLang"
	fmt.Println("cc = ", cc, "dd = ", dd)

	// 多行声明多个变量
	var (
		ee int  = 400
		ff bool = true
		gg      = false
	)
	fmt.Println("ee = ", ee, "ff = ", ff)
	fmt.Println("gg = ", gg)
}

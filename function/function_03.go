package main

import "fmt"

func main() {
	//fmt.Println("Hello World!")

	f1r1 := foo1(100, 200)
	f2r1, f2r2 := foo2(200, "Hello World!")
	f3r1, f3r2 := foo3(300, true)
	f4r1, f4r2 := foo4(400, "Hello C++")

	fmt.Println("f1r1 =", f1r1)
	fmt.Println("f2r1 =", f2r1, "f2r2 =", f2r2)
	fmt.Println("f3r1 =", f3r1, "f3r2 =", f3r2)
	fmt.Println("f4r1 =", f4r1, "f4r2 =", f4r2)
}

func foo1(aa int, bb int) int {
	println("aa =", aa, "bb =", bb)
	cc := 100
	return cc
}

// 如果返回值一样可以简写
func foo2(aa int, bb string) (cc int, dd int /*cc, dd int*/) {
	println("aa =", aa, "bb =", bb)
	return 100, 200
}

func foo3(aa int, bb bool) (int, string) {
	println("aa =", aa, "bb =", bb)
	return 100, "Hello GoLang"
}

// cc, dd为局部变量
func foo4(aa int, bb string) (cc int, dd string) {
	println("aa =", aa, "bb =", bb)
	cc = 200
	dd = "Hello World"
	return
}

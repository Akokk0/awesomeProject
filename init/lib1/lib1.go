package lib1

import "fmt"

// 函数使用大驼峰命名可以向外部暴露该函数，如果小写开头只能包自己内部使用
func Lib1Test() {
	fmt.Println("Lib1Test() ...")
}

func init() {
	fmt.Println("lib1 init() ...")
}

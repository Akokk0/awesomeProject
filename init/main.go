package main

import (
	// 匿名导包
	_ "main/init/lib1"

	// 起别名
	// myLib2 "main/init/lib2"

	// 全局导入，相当于JS里的import * from xxx
	. "main/init/lib2"
)

func main() {

	// 导包不使用加_
	// lib1.Lib1Test()

	// 导包起别名
	// myLib2.Lib2Test()

	// 使用.导包可以直接使用所有函数
	Lib2Test()
}

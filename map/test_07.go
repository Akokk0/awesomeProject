package main

import "fmt"

func main() {
	/*var map1 map[string]string
	if map1 == nil {
		fmt.Println("map为空")
	}*/

	// 第一种定义方法
	var map1 map[string]string
	// 如果不确定开辟空间的大小可以不写
	map1 = make(map[string]string)
	// 一次性赋值
	/*map1 = map[string]string{
		"one":   "C++",
		"two":   "Java",
		"There": "GO",
	}*/
	map1["one"] = "C++"
	map1["two"] = "Java"
	map1["There"] = "Go"
	fmt.Println(map1)

	// 第二种定义方法
	map2 := make(map[int]string)
	map2[1] = "C++"
	map2[2] = "Java"
	map2[3] = "Go"
	fmt.Println(map2)

	// 第三种定义方法
	map3 := map[string]string{
		"one":   "C++",
		"two":   "Java",
		"there": "Go",
	}
	fmt.Println(map3)

	// 以下不常用
	var map4 map[string]string = map[string]string{
		"one":   "C++",
		"two":   "Java",
		"there": "Go",
	}
	fmt.Println(map4)
	var map5 = map[string]string{
		"one":   "C++",
		"two":   "Java",
		"there": "Go",
	}
	fmt.Println(map5)

}

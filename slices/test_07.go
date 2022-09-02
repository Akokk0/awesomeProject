package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	// [0, 2) 左闭右开
	s1 := s[0:2]
	fmt.Println(s)
	fmt.Println(s1)
	// [0 : ] 表示全部取值 [ : 最后一个索引值] 表示最后一个不取值

	// 该处赋值还是使用引用赋值
	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1)

	s2 := make([]int, 3)
	// 使用copy会进行深拷贝
	copy(s2, s)
	fmt.Println("=============================")

	fmt.Println(s)
	fmt.Println(s2)

	s2[0] = 200

	fmt.Println(s)
	fmt.Println(s2)

}

package main

import "fmt"

func main() {
	/*// 固定长度数组
	var myArray [10]int
	// 定义之后赋值
	myArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	myArray2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for i := 0; i < len(myArray2); i++ {
		fmt.Println(myArray2[i])
	}

	// 增强for循环
	for index, value := range myArray2 {
		fmt.Println("index =", index, "value =", value)
	}*/

	// 声明slice的第一种方式
	// slice1 := []int{1, 2, 3, 4, 5} // slice len = 5, slice = [1 2 3 4 5]

	// 声明slice的第二种方式，声明了slice但是没有分配空间
	// var slice1 []int
	// 使用前需要给slice分配空间
	// slice1 = make([]int, 3) // slice len = 3, slice = [0 0 0]

	//声明slice的第三种方式，声明slice并开辟新的空间
	// var slice1 []int = make([]int, 3) // slice len = 3, slice = [0 0 0]

	// 声明slice的第四种方式，使用:=简写
	slice1 := make([]int, 3) // slice len = 3, slice = [0 0 0]

	fmt.Printf("slice len = %d, slice = %v", len(slice1), slice1)

}

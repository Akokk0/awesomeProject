package main

import "fmt"

func printMap(cMap map[string]string) {
	// 这里传递过来的map是引用map
	for index, value := range cMap {
		fmt.Println("index =", index, " value =", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["China"] = "ChengDu"
}

func main() {
	/*cityMap := map[string]string{
		"China": "ShangHai",
		"Japan": "Tokyo",
		"USA":   "NewYork",
	}

	printMap(cityMap)

	fmt.Println("============================")

	// 删除一个键
	delete(cityMap, "China")
	// 修改一个键的值
	cityMap["USA"] = "Los Angle"

	printMap(cityMap)*/

	// map不论空间大小是否确定传给函数都是引用传递
	cityMap := make(map[string]string, 3)
	cityMap["China"] = "ShangHai"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "Los Angle"

	fmt.Println(cityMap)

	fmt.Printf("=================================")

	changeValue(cityMap)
	fmt.Println(cityMap)

}

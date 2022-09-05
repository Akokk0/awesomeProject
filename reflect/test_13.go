// 反射基础用法
/*package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type : ", reflect.TypeOf(arg))
	fmt.Println("type : ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.23456
	reflectNum(num)
}*/

// 反射进阶用法
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

// 这里方法必须要写为u User而不是u *User，不然不能被NumMethod检测到
func (u User) Call() {
	fmt.Println("User is called...")
	fmt.Printf("%v\n", u)
}

func main() {
	user := User{
		Id:   1,
		Name: "张三",
		Age:  18,
	}

	DoFieldAndMethod(user)
}

func DoFieldAndMethod(input interface{}) {
	// 获取input的Type
	inputType := reflect.TypeOf(input)
	// fmt.Println("inputType is :", inputType.Name())
	// 获取input的Value
	inputValue := reflect.ValueOf(input)
	// fmt.Println("inputValue is :", inputValue)

	/**
	通过type获取里面的字段
	1、获取interface的reflect.Type，通过Type得到NumField，进行遍历
	2、得到每个field，数据类型
	3、通过field有一个Interface()方法得到对应的value
	*/
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过type获取里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

}

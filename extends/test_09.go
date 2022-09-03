package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Walk() {
	fmt.Println("Walk()...")
}

func (h *Human) Run() {
	fmt.Println("Run...")
}

type SuperMan struct {
	// 继承Human类
	Human
	// 自己的字段
	Level int
}

func (s *SuperMan) SuperHeroPunch() {
	fmt.Println("SuperHeroPunch()...")
}

func main() {
	// 父类
	hum := Human{
		Name: "张三",
		Age:  18,
	}
	hum.Walk()
	hum.Run()

	// 子类定义
	super := SuperMan{
		Human: Human{Name: "张三", Age: 18},
		Level: 200,
	}

	super2 := SuperMan{Human{
		"张三",
		18,
	}, 200}

	fmt.Println(super)
	fmt.Println(super2)

	//========================

	var super3 SuperMan
	super3.Name = "李四"
	super3.Age = 20
	super3.Level = 200

	fmt.Println(super3)

}

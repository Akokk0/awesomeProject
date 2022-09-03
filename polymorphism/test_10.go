package main

import (
	"fmt"
)

// interface实际上是一个指针
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (c *Cat) Sleep() {
	fmt.Println("Cat is sleeping...")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "Cat"
}

func main() {
	cat := Cat{color: "Pink"}
	cat.Sleep()
	printAnimal(&cat)
}

func printAnimal(animal Animal) {
	fmt.Printf("Animal is %s\n", animal.GetType())
	fmt.Printf("Animal's color is %s\n", animal.GetColor())
}

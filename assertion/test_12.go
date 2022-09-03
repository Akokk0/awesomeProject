package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Optional interface {
	OptionalFunc()
}

// 同时继承ReadBook接口和WriteBook接口
type Book struct{}

func (b *Book) ReadBook() {
	fmt.Println("ReadBook()...")
}

func (b *Book) WriteBook() {
	fmt.Println("WriteBook()...")
}

func main() {
	bp := &Book{}
	var r Reader = bp

	// tr, _ := r.(Optional) // 类型断言相当于如果不是该类型value则为nil，如果value有值相当于确实为该类型，所以可以当作类型转换来用
	// fmt.Printf("type is %T\n", tr)

	r.ReadBook()

	var w Writer
	// 这里不使用类型断言也没问题
	w = r.(Writer)

	w.WriteBook()
}

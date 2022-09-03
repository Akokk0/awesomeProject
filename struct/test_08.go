package main

import "fmt"

/*type myInt int*/

type Book struct {
	title  string
	author string
}

// 定义类
type Person struct {
	// 变量名首字母大写代表public，小写代表private，在本包都是可以访问的
	name string
	age  int
}

// 包括方法函数也是
func (p *Person /*这个的意思是绑定为类Person的方法*/) GetName() string {
	return p.name
}

func (p *Person) SetName(newName string) {
	p.name = newName
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) SetAge(newAge int) {
	p.age = newAge
}

func main() {

	/*var number myInt = 10

	fmt.Printf("number is %T\n", number) // number is main.myInt
	fmt.Printf("number = %v\n", number) // number = 10*/

	/*// 第一种定义方式
	var book1 Book = Book{
		title:  "GoLang",
		author: "ZhangSan",
	}

	// 第二种定义方式
	var book2 Book
	book2.title = "Java"
	book2.author = "LiSi"

	// 第三种定义方式
	book3 := Book{
		title:  "C++",
		author: "WangWu",
	}

	fmt.Println("Book1 is ", book1)
	fmt.Println("Book2 is ", book2)
	fmt.Println("Book3 is ", book3)*/

	/*book := Book{
		title:  "GoLang",
		author: "ZhangSan",
	}

	fmt.Println(book)

	changeAuthor(book)

	fmt.Println(book)

	changeOrigin(&book)

	fmt.Println(book)*/

	// 本包中都是可以访问的
	/*var person Person
	person.age = 20
	person.name = "张三"

	fmt.Println(person)

	person.SetName("李四")
	person.SetAge(18)

	fmt.Println(person)

	fmt.Println(person.GetName(), person.GetAge())*/

	person := Person{
		name: "王五",
		age:  30,
	}

	fmt.Println(person.GetName(), person.GetAge())

}

// 这里证明struct是值传递，其实所有go的数据类型都是值传递，只不过slice和map是内部带有指针字段所有产生引用传递的效果
func changeAuthor(book Book) {
	book.author = "Akokko"
}

func changeOrigin(book *Book) {
	book.author = "Akokko"
}

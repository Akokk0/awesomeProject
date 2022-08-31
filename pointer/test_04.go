package main

import "fmt"

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	a := 100
	b := 200

	fmt.Println("a =", a, "b =", b)

	// swap
	swap(&a, &b)

	fmt.Println("a =", a, "b =", b)
}

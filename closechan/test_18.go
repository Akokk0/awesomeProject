package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		fmt.Println("go routine start...")
		defer fmt.Println("go routine end...")
		for i := 0; i < 5; i++ {
			c <- i
		}

		// 未关闭channel fatal error: all goroutines are asleep - deadlock!
		close(c)
	}()

	/*for {
		if data, ok := <-c; ok {
			fmt.Println("Value is", data)
		} else {
			break
		}
	}*/

	// 👆代码可简写为
	for data := range c {
		println("value is", data)
	}

	fmt.Println("main routine end...")
}

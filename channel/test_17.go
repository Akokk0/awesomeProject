/*package main

import "fmt"

func main() {
	// 声明管道
	c := make(chan int)

	go func() {
		fmt.Println("go routine正在执行...")
		defer fmt.Println("go routine执行结束...")
		// 这里如果先运行会等到取值的地方再继续
		c <- 10
	}()

	// 这里如果先运行会等到赋值的地方再继续
	num := <-c

	fmt.Printf("num is %d\n", num)
	fmt.Println("main routine end...")
}*/

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // 带有缓冲的管道
	fmt.Println("len(c) =", len(c), "cap(c) =", cap(c))

	go func() {
		defer fmt.Println("go routine运行结束...")
		fmt.Println("go routine正在运行...")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("len(c) =", len(c), "cap(c) =", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num is", num, "len(c) =", len(c), "cap(c) =", cap(c))
	}

	fmt.Println("main routine结束...")

}

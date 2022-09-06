/*package main

import (
	"fmt"
	"time"
)

// 从Goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("New Goroutine is %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主Goroutine
func main() {
	// 如果主goroutine死亡，其他goroutine也会随之死亡

	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("Main Goroutine is %d\n", i)
		time.Sleep(time.Second)
	}
}*/

package main

import (
	"fmt"
	"time"
)

func main() {

	/*go func() {
		defer println("A.defer")
		// return 这里返回后面都不会执行
		func() {
			defer println("B.defer")
			runtime.Goexit() // 要想在这里结束进程不能使用return
			fmt.Println("B")
		}()
	}()*/

	// 使用go关键字开启的进程不能使用传统方法接收返回值
	go func(a int, b int) bool {
		fmt.Printf("a is %d, b is %d\n", a, b)
		return true
	}(10, 20)

	for {
		time.Sleep(time.Second)
	}

}

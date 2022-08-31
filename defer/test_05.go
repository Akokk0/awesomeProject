package main

// return比defer先执行
func deferFunc() int {
	// 有多个defer则先先进后出，与栈顺序相同
	defer println("deferFunc1()...")
	defer println("deferFunc2()...")
	defer println("deferFunc3()...")
	return 0
}

func returnFunc() int {
	println("returnFunc()...")
	return 0
}

func deferAndReturn() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	a := deferAndReturn()
	println(a)
}

package main

import "fmt"

func startCoroutines(i int32) {
	fmt.Println("Good,start coroutines successfully! i=>", i)
}

func main() {
	//go 关键字带函数 开启协程
	go func() {
		fmt.Println("hello world")
	}()

	fmt.Println("golang 开启协程例子")
	for i := int32(0); i < 10; i++ {
		go startCoroutines(i)
	}
	//等待任意键输入，退出程序
	fmt.Scanln()
}

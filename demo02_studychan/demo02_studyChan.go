package main

import "fmt"

func main() {
	//通道的定义 var 变量名 chan 类型
	var ch1 chan int
	//读、写前，要实例化，没实例化的chan，读、写会panic
	ch1 = make(chan int)

	//也可以直接这样定义
	//比较简洁，也顺便实例化
	ch2 := make(chan int)

	//如果没初始化 运行输出：fatal error: all goroutines are asleep - deadlock!
	//因为写入通道的时候，没有接收方
	//要开启一个协程去写入
	go func() {
		ch1 <- 8
		ch2 <- 9
	}()
	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
	//运行输出 8 9
	//最后记得关闭 channel
	close(ch1)
	close(ch2)
}

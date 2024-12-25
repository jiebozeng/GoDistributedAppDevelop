package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var initVar int64

func doInit() {
	//初始化代码
	initVar = 99
}

func main() {
	var i int64 = 0
	var mu sync.Mutex
	//开始锁
	mu.Lock()
	//对共享资源进行操作
	i++
	//解锁
	mu.Unlock()

	var rwMu sync.RWMutex
	//读锁，允许多个协程同时读
	rwMu.RLock()
	//只读操作
	//这里读取i，输出到控制台
	fmt.Println(i)
	rwMu.RUnlock()

	//写锁，只允许一个协程写
	rwMu.Lock()
	//写入操作
	i++
	//解锁
	rwMu.Unlock()

	once.Do(doInit)
}

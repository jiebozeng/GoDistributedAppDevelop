package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// 定义常量
const (
	workerBits  uint8 = 10                      // 工作节点ID占用的位数
	numberBits  uint8 = 12                      // 序列号占用的位数
	workerMax   int64 = -1 ^ (-1 << workerBits) // 工作节点ID的最大值
	numberMax   int64 = -1 ^ (-1 << numberBits) // 序列号的最大值
	timeShift   uint8 = workerBits + numberBits // 时间戳左移的位数
	workerShift uint8 = numberBits              // 工作节点ID左移的位数
	startTime   int64 = 1525705533000           // 起始时间戳，用于计算相对时间
)

// Worker 结构体定义
type Worker struct {
	mu        sync.Mutex // 互斥锁，保证并发安全
	timestamp int64      // 上次生成ID的时间戳
	workerId  int64      // 工作节点ID
	number    int64      // 序列号
}

// NewWorker 创建一个新的工作节点实例
func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity") // 工作节点ID超出范围
	}
	// 初始化工作节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

// GetId 生成并返回一个唯一的ID
func (w *Worker) GetId() int64 {
	w.mu.Lock()         // 加锁，保证并发安全
	defer w.mu.Unlock() // 解锁

	now := time.Now().UnixNano() / 1e6 // 获取当前时间戳（毫秒）

	// 如果当前时间戳与上次生成ID的时间戳相同，则序列号加1
	if w.timestamp == now {
		w.number++
		// 如果序列号超过最大值，则等待下一毫秒
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// 如果当前时间戳与上次生成ID的时间戳不同，则重置序列号
		w.number = 0
		w.timestamp = now
	}

	// 生成ID：时间戳左移timeShift位 | 工作节点ID左移workerShift位 | 序列号
	ID := int64((now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
	return ID
}

func main() {
	// 创建一个新的工作节点实例
	node, err := NewWorker(1)
	if err != nil {
		panic(err) // 如果创建失败，则抛出异常
	}

	// 无限循环生成并打印ID
	i := 0
	for {
		fmt.Println(node.GetId())
		i++
		//生成50个
		if i > 50 {
			break
		}
	}
}

package main

import (
	"fmt"
	"go-design-patterns/01_singleton/singleton"
	"sync"
)

// 全局goroutine数量定义
const goroutineCount = 100

// 客户端【使用方】
func main() {
	//serial()
	parallel()
}

// 单个goroutine
func serial() {
	instance1 := singleton.GetInstance()
	instance2 := singleton.GetInstance()
	if instance1 != instance2 {
		fmt.Println("instance1不等于instance2")
	} else {
		fmt.Println("instance1等于instance2")
	}
}

// 多个goroutine
func parallel() {
	// 同步等待组
	wg := sync.WaitGroup{}
	wg.Add(goroutineCount)
	// 定义数组存储
	singletons := [goroutineCount]*singleton.Singleton{}
	for i := 0; i < goroutineCount; i++ {
		// go的方式将数组内容填充
		go func(index int) {
			singletons[index] = singleton.GetInstance()
			wg.Done() // 同步等待组的计时器-1
		}(i)
	}
	wg.Wait() // 主goroutine进入阻塞状态
	for i := 1; i < goroutineCount; i++ {
		// 不等于nil可以进行操作
		if singletons[i] != singletons[i-1] {
			fmt.Printf("singletons%d不等于singletons%d\n", i, i-1)
		} else {
			fmt.Printf("singletons%v等于singletons%v\n", i, i-1)
		}
	}
}

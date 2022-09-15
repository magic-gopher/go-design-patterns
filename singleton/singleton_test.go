package singleton

import (
	"sync"
	"testing"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	if instance1 != instance2 {
		t.Fatal("实例化结构体不相等")
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*Singleton{}
	for i := 0; i < parCount; i++ {
		/*
			各种格式是匿名函数的写法
			go func(参数列表) {
				// 代码逻辑体
			}(实际入参数)
		*/
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait() // 让main goroutine进入阻塞状态
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("实例化结构体不相等")
		}
	}
}

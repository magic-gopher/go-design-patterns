package singleton

import "sync"

var lock = &sync.Mutex{}

// Singleton 结构体
type Singleton struct {
}

// singleton 结构体的指针全局变量
var instance *Singleton

// GetInstance 获取singleton结构体指针变量的函数
func GetInstance() *Singleton {
	if instance != nil {
		return instance
	} else {
		lock.Lock()         // 上锁
		defer lock.Unlock() // 释放锁
		if instance == nil {
			instance = &Singleton{}
		}
		return instance
	}
}

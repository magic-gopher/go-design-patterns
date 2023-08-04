package singleton

import "sync"

var (
	// singleton指针全局变量
	instance *singleton
	// 使函数只执行一次
	once sync.Once
)

// singleton 结构体
type singleton struct {
	value string
}

// GetInstance 获取singleton结构体指针函数
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{value: "singleton design"}
	})
	return instance
}

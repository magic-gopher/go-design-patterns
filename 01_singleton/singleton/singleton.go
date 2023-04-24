package singleton

import "sync"

var (
	// instance 单例实例变量
	instance *singleton
	// 使函数只执行一次
	once sync.Once
)

// singleton 单例结构体
type singleton struct {
}

// GetInstance 获取实例函数
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

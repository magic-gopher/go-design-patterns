package factory_method

import "fmt"

// 工厂生产的具体产品

type Coffee interface {
	String()
}

// mochaCoffee 摩卡咖啡
type mochaCoffee struct {
}

// NewMochaCoffee 获取mochaCoffee指针
func NewMochaCoffee() *mochaCoffee {
	return &mochaCoffee{}
}

// String 实现Coffee接口
func (mc *mochaCoffee) String() {
	fmt.Println("mocha coffee")
}

// latteCoffee 拿铁咖啡
type latteCoffee struct {
}

// NewLatteCoffee 获取latteCoffee指针
func NewLatteCoffee() *latteCoffee {
	return &latteCoffee{}
}

// String 实现Coffee接口
func (lc *latteCoffee) String() {
	fmt.Println("latte coffee")
}

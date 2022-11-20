package coffee

import "fmt"

// Coffee 咖啡接口
type Coffee interface {
	Info() string // 咖啡详情
}

// LatteCoffee 拿铁咖啡
type LatteCoffee struct {
}

// AmericanCoffee 美式咖啡
type AmericanCoffee struct {
}

// Info AmericanCoffee结构体的方法
func (a *AmericanCoffee) Info() string {
	return fmt.Sprintln("美式咖啡")
}

// Info LatteCoffee结构体的方法
func (l *LatteCoffee) Info() string {
	return fmt.Sprintln("拿铁咖啡")
}

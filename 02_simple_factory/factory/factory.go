package factory

import "go-design-patterns/02_simple_factory/coffee"

// CreateCoffee 制作咖啡
func CreateCoffee(coffeeType int) coffee.Coffee {
	if coffeeType == 1 {
		// 类型1就制作拿铁咖啡
		return &coffee.LatteCoffee{}
	} else {
		// 类型2就制作美式咖啡
		return &coffee.AmericanCoffee{}
	}
}

package factory

import "go-design-patterns/02_simple_factory/coffee"

// CoffeeFactory 咖啡工厂生产Coffee
type CoffeeFactory interface {
	CreateCoffee() coffee.Coffee
}

// LatteCoffeeFactory 拿铁咖啡工厂
type LatteCoffeeFactory struct {
}

func NewCoffeeFactory(coffeeType string) CoffeeFactory {
	if coffeeType == "latte" {
		return &LatteCoffeeFactory{}
	} else if coffeeType == "american" {
		return &AmericanCoffeeFactory{}
	} else {
		panic("coffee type is defined")
	}
}

// AmericanCoffeeFactory 美式咖啡工厂
type AmericanCoffeeFactory struct {
}

// CreateCoffee 生产拿铁咖啡
func (factory *LatteCoffeeFactory) CreateCoffee() coffee.Coffee {
	return &coffee.LatteCoffee{}
}

// CreateCoffee 生产美式咖啡
func (factory *AmericanCoffeeFactory) CreateCoffee() coffee.Coffee {
	return &coffee.AmericanCoffee{}
}

package simple_factory_test

import (
	"go-design-patterns/02_simple_factory/simple_factory"
	"testing"
)

// 制作拿铁咖啡
func TestCreateLatteCoffee(t *testing.T) {
	// 创建咖啡工厂实例
	coffeeFactory := simple_factory.CoffeeFactory{}
	// 制作咖啡
	coffee := coffeeFactory.CreateCoffee("latte")
	// 验证
	if coffee.Info() == "latte" {
		t.Log("latte successfully made")
	}
}

// 制作美式咖啡
func TestAmericanoCoffee(t *testing.T) {
	// 创建咖啡工厂实例
	coffeeFactory := simple_factory.CoffeeFactory{}
	// 制作咖啡
	coffee := coffeeFactory.CreateCoffee("americano")
	// 验证
	if coffee.Info() == "americano" {
		t.Log("americano successfully made")
	}
}

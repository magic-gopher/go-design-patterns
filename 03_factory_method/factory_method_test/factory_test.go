package factory_method_test

import (
	"fmt"
	"go-design-patterns/03_factory_method/factory_method"
	"testing"
)

func TestCreateLatteCoffee(t *testing.T) {
	// 创建拿铁咖啡工厂
	latteCoffeeFactory := factory_method.LatteCoffeeFactory{}
	// 制作拿铁咖啡
	coffee := latteCoffeeFactory.CreateCoffee(true, 5)
	fmt.Println(coffee.Result())
}

func TestCreateAmericanoCoffee(t *testing.T) {
	// 创建美式咖啡工厂
	americanoCoffeeFactory := factory_method.AmericanoCoffeeFactory{}
	// 制作美式咖啡
	coffee := americanoCoffeeFactory.CreateCoffee(false, 4)
	fmt.Println(coffee.Result())
}

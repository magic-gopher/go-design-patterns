package main

import (
	"fmt"
	"go-design-patterns/03_factory_method/factory"
)

// 客户端【使用方】
func main() {
	// latte
	//coffeeFactory := factory.NewCoffeeFactory("latte")
	//coffee := coffeeFactory.CreateCoffee()
	//fmt.Println(coffee.Info())

	// american
	//coffeeFactory := factory.NewCoffeeFactory("american")
	//coffee := coffeeFactory.CreateCoffee()
	//fmt.Println(coffee.Info())

	// coffee type is defined
	coffeeFactory := factory.NewCoffeeFactory("cappuccino")
	coffee := coffeeFactory.CreateCoffee()
	fmt.Println(coffee.Info())
}

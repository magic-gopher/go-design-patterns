package main

import (
	"fmt"
	"go-design-patterns/02_simple_factory/factory"
)

// 客户端【使用方】
func main() {
	coffeeA := factory.CreateCoffee(1)
	fmt.Println(coffeeA.Info()) // 拿铁咖啡
	coffeeB := factory.CreateCoffee(2)
	fmt.Println(coffeeB.Info()) // 美式咖啡
}

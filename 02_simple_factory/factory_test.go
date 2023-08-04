package simple_factory

import (
	"testing"
)

// TestCreateCoffee 制作拿铁咖啡
func TestLatteCoffee(t *testing.T) {
	coffeeFactory := CoffeeFactory{}
	coffee := coffeeFactory.CreateCoffee("latte")
	coffee.SetSugar(7)
	coffee.SetIceCake(true)
	coffee.Info()
}

// TestAmericanoCoffee 制作美式咖啡
func TestAmericanoCoffee(t *testing.T) {
	coffeeFactory := CoffeeFactory{}
	coffee := coffeeFactory.CreateCoffee("americano")
	coffee.SetSugar(4)
	coffee.SetIceCake(false)
	coffee.Info()
}

// TestCreateCoffee
func TestCreateCoffee(t *testing.T) {
	coffeeFactory := CoffeeFactory{}
	coffee := coffeeFactory.CreateCoffee("coffee")
	coffee.SetSugar(4)
	coffee.SetIceCake(false)
	coffee.Info()
}

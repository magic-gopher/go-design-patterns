package factory_method

import (
	"testing"
)

// Test && Benchmark

// TestMakeLatteCoffee 制作拿铁咖啡
func TestMakeLatteCoffee(t *testing.T) {
	// 咖啡店
	coffeeStore := NewCoffeeStore()
	// 制作咖啡
	coffee := coffeeStore.SalesCoffee(1)
	coffee.String()
}

// TestMakeMochaCoffee 制作摩卡咖啡
func TestMakeMochaCoffee(t *testing.T) {
	// 咖啡店
	coffeeStore := NewCoffeeStore()
	// 制作咖啡
	coffee := coffeeStore.SalesCoffee(2)
	coffee.String()
}

// TestNotListedCoffee 未上架的咖啡
func TestNotListedCoffee(t *testing.T) {
	// 咖啡店
	coffeeStore := NewCoffeeStore()
	// 制作咖啡
	coffee := coffeeStore.SalesCoffee(3)
	coffee.String()
}

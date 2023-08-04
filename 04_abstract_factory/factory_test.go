package abstract_factory

import (
	"testing"
)

// Test && Benchmark

// 1.咖啡 2.奶茶 3.烤牛肉 4.提拉米苏

// TestCoffee 售卖咖啡
func TestCoffee(t *testing.T) {
	drinkStore := NewDrinkStore()
	drinkStore.SalesGood(1)
}

// TestTeaMilk 售卖奶茶
func TestTeaMilk(t *testing.T) {
	drinkStore := NewDrinkStore()
	drinkStore.SalesGood(2)
}

// TestMeat 售卖烤牛肉
func TestMeat(t *testing.T) {
	foodStore := NewFoodStore()
	foodStore.SalesGood(3)
}

// TestCake 售卖提拉米苏
func TestCake(t *testing.T) {
	foodStore := NewFoodStore()
	foodStore.SalesGood(4)
}

package decorator

import "testing"

func TestCoffeeDecorator(t *testing.T) {
	// 创建一个基础咖啡对象
	coffee := &ConcreteCoffee{}

	// 添加牛奶装饰
	coffeeWithMilk := &MilkDecorator{
		CoffeeDecorator: &CoffeeDecorator{
			Coffee: coffee,
		},
	}

	// 添加糖装饰
	coffeeWithMilkAndSugar := &SugarDecorator{
		CoffeeDecorator: &CoffeeDecorator{
			Coffee: coffeeWithMilk,
		},
	}

	// 验证咖啡描述和价格是否符合预期
	expectedDescription := "普通咖啡 + 牛奶 + 糖"
	expectedCost := 15.0

	description := coffeeWithMilkAndSugar.GetDescription()
	cost := coffeeWithMilkAndSugar.GetCost()

	if description != expectedDescription {
		t.Errorf("Expected description: %s, but got: %s", expectedDescription, description)
	}

	if cost != expectedCost {
		t.Errorf("Expected cost: %f, but got: %f", expectedCost, cost)
	}
}

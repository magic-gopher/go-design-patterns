package simple_factory

// 简单工厂
// 注意简单工厂不是设计模式，只是写代码的一种规范

// CoffeeFactory 咖啡工厂
type CoffeeFactory struct {
}

// CreateCoffee 制作咖啡
func (cf *CoffeeFactory) CreateCoffee(CafeType string) Coffee {
	switch CafeType {
	case "latte":
		return &LatteCoffee{}
	case "americano":
		return &AmericanoCoffee{}
	}
	return nil
}

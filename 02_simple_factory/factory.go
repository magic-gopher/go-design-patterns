package simple_factory

// 简单工厂
// 注意简单工厂不是设计模式，只是写代码的一种规范

// CoffeeFactory 咖啡工厂
type CoffeeFactory struct {
}

// CreateCoffee 制作咖啡
func (cf *CoffeeFactory) CreateCoffee(types string) Coffee {
	switch types {
	case "latte":
		return NewLatteCoffee()
	case "americano":
		return NewAmericanoCoffee()
	default:
		panic("咖啡类型不存在!")
	}
}

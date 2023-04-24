package simple_factory

// 工厂

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

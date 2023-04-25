package factory_method

// 工厂接口

// CoffeeFactory 咖啡工厂
type CoffeeFactory interface {
	// CreateCoffee 制作咖啡
	CreateCoffee(milk bool, sugar int) Coffee
}

// LatteCoffeeFactory 拿铁咖啡工厂
type LatteCoffeeFactory struct {
}

func (lcf *LatteCoffeeFactory) CreateCoffee(milk bool, sugar int) Coffee {
	latteCoffee := &LatteCoffee{}
	// 加材料
	latteCoffee.SetMilk(milk)
	latteCoffee.SetSugar(sugar)
	return latteCoffee
}

// AmericanoCoffeeFactory 美式咖啡工厂
type AmericanoCoffeeFactory struct {
}

func (acf *AmericanoCoffeeFactory) CreateCoffee(milk bool, sugar int) Coffee {
	americanoCoffee := &AmericanoCoffee{}
	// 加材料
	americanoCoffee.SetMilk(milk)
	americanoCoffee.SetSugar(sugar)
	return americanoCoffee
}

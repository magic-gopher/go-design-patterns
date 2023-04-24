package simple_factory

// 抽象产品接口

// Coffee 咖啡
type Coffee interface {
	Info() string
}

// LatteCoffee 拿铁咖啡
type LatteCoffee struct {
}

func (lc *LatteCoffee) Info() string {
	return "latte"
}

// AmericanoCoffee 美式咖啡
type AmericanoCoffee struct {
}

func (ac *AmericanoCoffee) Info() string {
	return "americano"
}

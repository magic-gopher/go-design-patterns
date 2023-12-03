package decorator

// 咖啡

// Coffee 接口定义了咖啡的基本行为
type Coffee interface {
	GetDescription() string
	GetCost() float64
}

// ConcreteCoffee 是具体的咖啡类
type ConcreteCoffee struct{}

func (c *ConcreteCoffee) GetDescription() string {
	return "普通咖啡"
}

func (c *ConcreteCoffee) GetCost() float64 {
	return 10.0
}

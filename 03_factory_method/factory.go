package factory_method

// 生产具体产品的工厂

// CoffeeFactory 咖啡工厂
type CoffeeFactory interface {
	// MakeCoffee 制作咖啡
	MakeCoffee() Coffee
}

// mochaCoffeeFactory 摩卡咖啡工厂
type mochaCoffeeFactory struct {
}

// NewMochaCoffeeFactory 获取mochaCoffeeFactory结构体指针
func NewMochaCoffeeFactory() *mochaCoffeeFactory {
	return &mochaCoffeeFactory{}
}

// MakeCoffee 实现CoffeeFactory接口
func (mcf *mochaCoffeeFactory) MakeCoffee() Coffee {
	return NewMochaCoffee()
}

// LatteCoffeeFactory 拿铁咖啡工厂
type latteCoffeeFactory struct {
}

// NewLatteCoffeeFactory 获取latteCoffeeFactory结构体指针
func NewLatteCoffeeFactory() *latteCoffeeFactory {
	return &latteCoffeeFactory{}
}

// MakeCoffee 实现CoffeeFactory接口
func (lcf *latteCoffeeFactory) MakeCoffee() Coffee {
	return NewLatteCoffee()
}

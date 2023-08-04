package factory_method

// 售卖的商店

// coffeeStore 咖啡店
type coffeeStore struct {
}

// NewCoffeeStore 获取coffeeStore结构体指针
func NewCoffeeStore() *coffeeStore {
	return &coffeeStore{}
}

// SalesCoffee 售卖咖啡
func (cs *coffeeStore) SalesCoffee(coffeeType int) Coffee {
	switch coffeeType {
	case 1:
		latteCoffee := NewLatteCoffeeFactory().MakeCoffee()
		return latteCoffee
	case 2:
		mochaCoffee := NewMochaCoffeeFactory().MakeCoffee()
		return mochaCoffee
	default:
		panic("该类型咖啡店里为上架，请重新选择!")
	}
}

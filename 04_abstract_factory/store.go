package abstract_factory

// 商店

// Store 商店接口
type Store interface {
	SalesGood(goodType int)
}

// DrinkStore 饮品商店
type DrinkStore struct {
}

// NewDrinkStore 获取DrinkStore结构体指针
func NewDrinkStore() *DrinkStore {
	return &DrinkStore{}
}

func (ds *DrinkStore) SalesGood(goodType int) {
	drinkConcreteFactory := NewDrinkConcreteFactory()
	switch goodType {
	case 1:
		latteCoffee := drinkConcreteFactory.MakeLatteCoffee()
		latteCoffee.Info()
	case 2:
		englishTeaMilk := drinkConcreteFactory.MakeEnglishTeaMilk()
		englishTeaMilk.Info()
	}
}

// FoodStore 食品商店
type FoodStore struct {
}

// NewFoodStore 获取NewFoodStore结构体指针
func NewFoodStore() *FoodStore {
	return &FoodStore{}
}

func (fs *FoodStore) SalesGood(goodType int) {
	foodConcreteFactory := NewFoodConcreteFactory()
	switch goodType {
	case 3:
		roastBeef := foodConcreteFactory.MakeRoastBeef()
		roastBeef.Info()
	case 4:
		tiramisu := foodConcreteFactory.MakeTiramisu()
		tiramisu.Info()
	}
}

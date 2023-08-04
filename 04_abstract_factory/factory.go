package abstract_factory

// 工厂

// drinkAbstractFactory 饮品抽象工厂
type drinkAbstractFactory interface {
	// MakeLatteCoffee 制作拿铁咖啡
	MakeLatteCoffee() Coffee
	// MakeEnglishTeaMilk 制作英式奶茶
	MakeEnglishTeaMilk() TeaMilk
}

// foodAbstractFactory 食品抽象工厂
type foodAbstractFactory interface {
	// MakeRoastBeef 制作烤牛肉
	MakeRoastBeef() Meat
	// MakeTiramisu 制作提拉米苏
	MakeTiramisu() Cake
}

// drinkConcreteFactory 饮品具体工厂
type drinkConcreteFactory struct {
}

// NewDrinkConcreteFactory 获取drinkConcreteFactory结构体指针
func NewDrinkConcreteFactory() *drinkConcreteFactory {
	return &drinkConcreteFactory{}
}

func (dcf *drinkConcreteFactory) MakeLatteCoffee() Coffee {
	return NewLatteCoffee()
}

func (dcf *drinkConcreteFactory) MakeEnglishTeaMilk() TeaMilk {
	return NewEnglishTeaMilk()
}

// foodConcreteFactory 食品工厂
type foodConcreteFactory struct {
}

// NewFoodConcreteFactory 获取foodConcreteFactory结构体指针
func NewFoodConcreteFactory() *foodConcreteFactory {
	return &foodConcreteFactory{}
}

func (fcf *foodConcreteFactory) MakeRoastBeef() Meat {
	return NewRoastBeef()
}

func (fcf *foodConcreteFactory) MakeTiramisu() Cake {
	return NewTiramisu()
}

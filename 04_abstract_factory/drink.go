package abstract_factory

import "fmt"

// 饮品类

// Drink 饮品类接口
type Drink interface {
	Info()
}

// Coffee 咖啡类接口
type Coffee interface {
	Drink
}

// TeaMilk 奶茶类接口
type TeaMilk interface {
	Drink
}

// latteCoffee 拿铁咖啡
type latteCoffee struct {
}

// NewLatteCoffee 创建latteCoffee指针
func NewLatteCoffee() *latteCoffee {
	return &latteCoffee{}
}

func (lc *latteCoffee) Info() {
	fmt.Println("拿铁咖啡")
}

// englishTeaMilk 英式奶茶
type englishTeaMilk struct {
}

// NewEnglishTeaMilk 创建NewEnglishTeaMilk指针
func NewEnglishTeaMilk() *englishTeaMilk {
	return &englishTeaMilk{}
}

func (etm *englishTeaMilk) Info() {
	fmt.Println("英式奶茶")
}

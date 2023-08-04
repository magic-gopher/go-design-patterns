package simple_factory

import "fmt"

// 工厂生产的产品

// Coffee 咖啡接口
type Coffee interface {
	// Info 信息
	Info()
	// SetIceCake 加冰块
	SetIceCake(flag bool)
	// SetSugar 几分糖
	SetSugar(weight int)
}

// BaseCoffee 咖啡基础
type baseCoffee struct {
	// 冰块
	iceCake bool
	// 糖份
	sugar int
}

// latteCoffee 拿铁咖啡
type latteCoffee struct {
	base baseCoffee
}

func NewLatteCoffee() *latteCoffee {
	return &latteCoffee{}
}

func (lc *latteCoffee) Info() {
	fmt.Printf("拿铁咖啡:{%d分糖, 加冰块:%t}\n", lc.base.sugar, lc.base.iceCake)
}

func (lc *latteCoffee) SetIceCake(flag bool) {
	lc.base.iceCake = flag
}

func (lc *latteCoffee) SetSugar(weight int) {
	lc.base.sugar = weight
}

// americanoCoffee 美式咖啡
type americanoCoffee struct {
	base baseCoffee
}

func NewAmericanoCoffee() *americanoCoffee {
	return &americanoCoffee{}
}

func (ac *americanoCoffee) Info() {
	fmt.Printf("美式咖啡:{%d分糖, 加冰块:%t}\n", ac.base.sugar, ac.base.iceCake)
}

func (ac *americanoCoffee) SetIceCake(flag bool) {
	ac.base.iceCake = flag
}

func (ac *americanoCoffee) SetSugar(weight int) {
	ac.base.sugar = weight
}

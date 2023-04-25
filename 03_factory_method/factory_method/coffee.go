package factory_method

import "fmt"

// Coffee 咖啡公共的封装
type Coffee interface {
	// SetMilk 是否添加牛奶
	SetMilk(bool)
	// SetSugar 几分糖
	SetSugar(int)
	// Result 返回信息
	Result() string
}

// LatteCoffee 拿铁咖啡
type LatteCoffee struct {
	milk  bool
	sugar int
}

func (lc *LatteCoffee) SetMilk(milk bool) {
	lc.milk = milk
}

func (lc *LatteCoffee) SetSugar(sugar int) {
	lc.sugar = sugar
}

func (lc *LatteCoffee) Result() string {
	latteString := fmt.Sprintf("latte={milk : %t\n, sugar : %d\n}", lc.milk, lc.sugar)
	return latteString
}

// AmericanoCoffee 美式咖啡
type AmericanoCoffee struct {
	milk  bool
	sugar int
}

func (ac *AmericanoCoffee) SetMilk(milk bool) {
	ac.milk = milk
}

func (ac *AmericanoCoffee) SetSugar(sugar int) {
	ac.sugar = sugar
}

func (ac *AmericanoCoffee) Result() string {
	latteString := fmt.Sprintf("americano={milk : %t\n, sugar : %d\n}", ac.milk, ac.sugar)
	return latteString
}

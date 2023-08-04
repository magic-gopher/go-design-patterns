package abstract_factory

import "fmt"

// 食品类

// Food 食品类接口
type Food interface {
	Info()
}

// Meat 肉类接口
type Meat interface {
	Food
}

// Cake 糕点接口
type Cake interface {
	Food
}

// roastBeef 烤牛肉
type roastBeef struct {
}

// NewRoastBeef 创建roastBeef指针
func NewRoastBeef() *roastBeef {
	return &roastBeef{}
}

func (b *roastBeef) Info() {
	fmt.Println("烤牛肉")
}

type tiramisu struct {
}

// NewTiramisu 创建tiramisu指针
func NewTiramisu() *tiramisu {
	return &tiramisu{}
}

func (t tiramisu) Info() {
	fmt.Println("提拉米苏")
}

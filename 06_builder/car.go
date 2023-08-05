package builder

import "fmt"

// ICar 汽车产品接口
type ICar interface {
	// CarInfo 显示车辆信息
	CarInfo()
}

// BaseCar 车基本定义
type BaseCar struct {
	// 座位数
	seats int
	// 车轮
	wheel int
	// 颜色
	color string
	// 引擎
	engine string
}

// SportsCar 跑车
type SportsCar struct {
	base BaseCar
}

func (sc *SportsCar) CarInfo() {
	fmt.Printf("跑车:{座位数:%d, 车轮数:%d, 颜色:%s, 引擎:%s}\n", sc.base.seats, sc.base.wheel, sc.base.color, sc.base.engine)
}

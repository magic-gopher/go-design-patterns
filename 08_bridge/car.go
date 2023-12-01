package bridge

import "fmt"

// 车辆

// Vehicle 车辆抽象
type Vehicle interface {
	// Start 启动引擎函数
	Start()
	// SetEngine 设置引擎
	SetEngine(engine Engine)
}

// Car 轿车类型
type Car struct {
	// engine 引擎抽象接口
	engine Engine
}

func (c *Car) Start() {
	fmt.Println("轿车: 启动中...")
	c.engine.Start()
}

func (c *Car) SetEngine(engine Engine) {
	c.engine = engine
}

// Truck 卡车类型
type Truck struct {
	// engine 引擎抽象接口
	engine Engine
}

func (t *Truck) Start() {
	fmt.Println("卡车: 启动中...")
	t.engine.Start()
}

func (t *Truck) SetEngine(engine Engine) {
	t.engine = engine
}

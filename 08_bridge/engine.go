package bridge

import "fmt"

// 引擎

// Engine 引擎接口
type Engine interface {
	// Start 启动引擎函数
	Start()
}

// GasolineEngine 燃油引擎
type GasolineEngine struct {
}

func (ge *GasolineEngine) Start() {
	fmt.Println("燃油引擎: 启动中...")
}

// ElectricEngine 电能引擎
type ElectricEngine struct {
}

func (ee *ElectricEngine) Start() {
	fmt.Println("电能引擎: 启动中...")
}

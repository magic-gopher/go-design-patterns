package bridge

import "testing"

// Test && Benchmark

func TestVehicle(t *testing.T) {
	// 创建两种不同的引擎
	gasolineEngine := &GasolineEngine{}
	electricEngine := &ElectricEngine{}
	// 轿车类型
	car := &Car{}
	car.SetEngine(gasolineEngine)
	car.Start()
	// 卡车类型
	truck := &Truck{}
	truck.SetEngine(electricEngine)
	truck.Start()
}

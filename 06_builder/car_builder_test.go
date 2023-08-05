package builder

import "testing"

// Test && Benchmark

// 建造跑车
func TestBuildCar(t *testing.T) {
	// 具体建造者
	builder := NewSportsCarBuilder()
	// 汽车组装工程师设置具体构建者
	engineer := NewCarBuildEngineer(builder)
	car := engineer.BuildCar(2, 4, "黑色", "V16")
	car.CarInfo()
}

package prototype

import (
	"fmt"
	"testing"
)

// Test && Benchmark

func TestClone(t *testing.T) {
	// 山羊
	p1 := NewGoat()
	p1.Eat("草")
	fmt.Printf("实例地址信息:{%p}, 实例信息:{%v}\n", &p1, p1)
	// 克隆
	p2 := p1.Clone()
	p2.Eat("青苔")
	fmt.Printf("实例地址信息:{%p}, 实例信息:{%v}\n", &p2, p2)
}

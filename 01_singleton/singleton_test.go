package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetInstance 单元测试
func TestGetInstance(t *testing.T) {
	assert.Equal(t, GetInstance(), GetInstance())
}

// BenchmarkGetInstance 性能测试
func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Error("not the same instance")
			}
		}
	})
}

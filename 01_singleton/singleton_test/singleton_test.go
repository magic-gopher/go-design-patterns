package singleton_test

import (
	"github.com/stretchr/testify/assert"
	"go-design-patterns/01_singleton/singleton"
	"testing"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, singleton.GetInstance(), singleton.GetInstance())
}

func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetInstance() != singleton.GetInstance() {
				b.Error("instance not equation")
			}
		}
	})
}

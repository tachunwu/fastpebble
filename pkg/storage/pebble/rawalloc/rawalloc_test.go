package rawalloc

import (
	"fmt"
	"testing"
)

var sizes = []int{16, 100, 1024, 1024 * 10, 1024 * 100, 1024 * 1024}

func BenchmarkRawalloc(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("rawalloc-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = New(size, size)
			}
		})
	}
}

func BenchmarkMake(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("make-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = make([]byte, size)
			}
		})
	}
}

package youdao

import (
	"testing"
)

func BenchmarkSign(b *testing.B) {
	c := &Client{"hello", "world"}
	for i := 0; i < b.N; i++ {
		c.sign("你好", "3xhg3b")
	}
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randString(6)
	}
}

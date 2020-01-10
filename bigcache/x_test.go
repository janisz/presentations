package bigcache

import "testing"

var sink = 0

func BenchmarkMod(b *testing.B) {
        for n := 0; n < b.N; n++ {
                sink = n % 1024
        }
}

func BenchmarkBit(b *testing.B) {
        for n := 0; n < b.N; n++ {
                sink = n & 1023
        }
}

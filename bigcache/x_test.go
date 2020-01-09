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

func BenchmarkNotDefered(b *testing.B) {
        for n := 0; n < b.N; n++ {
                notDefered()
        }
}

func BenchmarkDefered(b *testing.B) {
        for n := 0; n < b.N; n++ {
                defered()
        }
}

func defered() {
    defer func() {}()
}

func notDefered() {
    func() {}()
}

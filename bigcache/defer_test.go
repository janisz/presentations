package bigcache

import "testing"

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

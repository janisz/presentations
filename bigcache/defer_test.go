package bigcache

import "testing"

var sink = 0

func BenchmarkNotDefered(b *testing.B) {
        for n := 0; n < b.N; n++ {
                notDefered()
        }
	b.Log(sink)
}

func BenchmarkDefered(b *testing.B) {
        for n := 0; n < b.N; n++ {
                defered()
        }
	b.Log(sink)
}

func defered() {
    defer func() {
	    sink++
    }()
}

func notDefered() {
    func() {
	    sink++
    }()
}

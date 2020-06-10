package main

import (
	"testing"
)

// BenchmarkConvertMapToList1-12           2000000000               0.26 ns/op            0 B/op          0 allocs/op
// BenchmarkConvertMapToList2-12           2000000000               1.80 ns/op            0 B/op          0 allocs/op
// BenchmarkIncrement1-12                  2000000000               0.52 ns/op            0 B/op          0 allocs/op
// BenchmarkIncrement2-12                  2000000000               0.26 ns/op            0 B/op          0 allocs/op
// BenchmarkIncrementP1-12                 2000000000               1.80 ns/op            0 B/op          0 allocs/op
// BenchmarkIncrementP2-12                 2000000000               1.80 ns/op            0 B/op          0 allocs/op

type Piyo struct {
	Counter int
}

func Increment(piyo Piyo) {
	piyo.Counter += 1
}

func IncrementP(piyo *Piyo) {
	piyo.Counter += 1
}

func BenchmarkConvertMapToList1(b *testing.B) {
	var piyo = Piyo{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		piyo.Counter += 1
	}
}

func BenchmarkConvertMapToList2(b *testing.B) {
	var ppiyo = &Piyo{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ppiyo.Counter += 1
	}
}

func BenchmarkIncrement1(b *testing.B) {
	var piyo = Piyo{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Increment(piyo)
	}
}

func BenchmarkIncrement2(b *testing.B) {
	var ppiyo = &Piyo{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Increment(*ppiyo)
	}
}

func BenchmarkIncrementP1(b *testing.B) {
	var piyo = Piyo{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IncrementP(&piyo)
	}
}

func BenchmarkIncrementP2(b *testing.B) {
	var ppiyo = &Piyo{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IncrementP(ppiyo)
	}
}

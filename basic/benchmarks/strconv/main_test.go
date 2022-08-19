package main

import (
	"strconv"
	"testing"
)

// $ go test -bench . ./int_vs_int64/int_vs_int64_test.go -benchmem
// goos: linux
// goarch: amd64
// BenchmarkAtoi-12                273517682                4.41 ns/op            0 B/op          0 allocs/op
// BenchmarkParseInt-12            98031213                12.1 ns/op             0 B/op          0 allocs/op
// BenchmarkItoa-12                518426824                2.30 ns/op            0 B/op          0 allocs/op
// BenchmarkFormatInt-12           504426722                2.29 ns/op            0 B/op          0 allocs/op
// PASS
// ok      command-line-arguments  6.436s

func BenchmarkAtoi(b *testing.B) {
	str := "1"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Atoi(str)
	}
}

func BenchmarkParseInt1(b *testing.B) {
	b.ResetTimer()
	str := "1"
	for i := 0; i < b.N; i++ {
		strconv.ParseInt(str, 10, 64)
	}
}

func BenchmarkParseInt2(b *testing.B) {
	b.ResetTimer()
	str := "1"
	for i := 0; i < b.N; i++ {
		num, _ := strconv.Atoi(str)
		_ = int64(num)
	}
}

func BenchmarkItoa(b *testing.B) {
	var num int = 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
}

func BenchmarkFormatInt64(b *testing.B) {
	var num int64 = 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
}

func BenchmarkFormatFloat(b *testing.B) {
	v := "3.1415926535"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseFloat(v, 64)
	}
}

func BenchmarkFormatIntFloat1(b *testing.B) {
	v := "3"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseFloat(v, 64)
	}
}

func BenchmarkFormatIntFloat2(b *testing.B) {
	v := "3"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		num, _ := strconv.Atoi(v)
		_ = float64(num)
	}
}

package handling_error

import (
	"fmt"
	"testing"
)

// Printfを使わなくて済むなら使わないほうが早い
// Printfは%vよりも明示的に型指定したほうが早い

// BenchmarkPrintf1-12     20000000                55.3 ns/op            16 B/op          1 allocs/op
// BenchmarkPrintf2-12     10000000               165 ns/op              32 B/op          2 allocs/op
// BenchmarkPrintf3-12     10000000               207 ns/op              16 B/op          1 allocs/op

func BenchmarkPrintf1(b *testing.B) {
	err := fmt.Errorf("Error")
	var tmp string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmp = "Error=" + err.Error()
	}
	_ = tmp
}

func BenchmarkPrintf2(b *testing.B) {
	err := fmt.Errorf("Error")
	var tmp string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmp = fmt.Sprintf("Error=%s", err.Error())
	}
	_ = tmp
}

func BenchmarkPrintf3(b *testing.B) {
	err := fmt.Errorf("Error")
	var tmp string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmp = fmt.Sprintf("Error=%v", err)
	}
	_ = tmp
}

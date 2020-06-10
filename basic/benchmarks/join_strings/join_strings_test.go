package join_strings

import (
	"fmt"
	"strings"
	"testing"
)

// BenchmarkAddStrings1-12         50000000                29.2 ns/op             0 B/op          0 allocs/op
// BenchmarkAddStrings2-12         20000000                54.2 ns/op             4 B/op          1 allocs/op
// BenchmarkAddStrings3-12         10000000               129 ns/op              12 B/op          3 allocs/op
// BenchmarkAddStrings4-12          5000000               353 ns/op              68 B/op          5 allocs/op

func BenchmarkAddStrings1(b *testing.B) {
	list := []string{"a", "b", "c", "d"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = list[0] + list[1] + list[2] + list[3]
	}
}

func BenchmarkAddStrings2(b *testing.B) {
	list := []string{"a", "b", "c", "d"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Join(list, "")
	}
}

func BenchmarkAddStrings3(b *testing.B) {
	list := []string{"a", "b", "c", "d"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		baseStr := ""
		for _, word := range list {
			baseStr += word
		}
	}
}

func BenchmarkAddStrings4(b *testing.B) {
	list := []string{"a", "b", "c", "d"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s%s", list[0], list[1], list[2], list[3])
	}
}

package main

import "testing"

type Piyo struct {
	Name string
}

var userMap = map[string]Piyo{
	"A": Piyo{Name: "AAAAAAAAAA"},
	"B": Piyo{Name: "AAAAAAAAAA"},
	"C": Piyo{Name: "AAAAAAAAAA"},
	"D": Piyo{Name: "AAAAAAAAAA"},
	"E": Piyo{Name: "AAAAAAAAAA"},
}

// BenchmarkConvertMapToList2-12            3000000               452 ns/op             240 B/op          2 allocs/op
// BenchmarkConvertMapToList1-12            3000000               616 ns/op             240 B/op          4 allocs/op

func BenchmarkConvertMapToList1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		users := make([]Piyo, len(userMap))
		for _, user := range userMap {
			users = append(users, user)
		}
	}
}

func BenchmarkConvertMapToList2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		users := []Piyo{}
		for _, user := range userMap {
			users = append(users, user)
		}
	}
}

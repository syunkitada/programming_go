package main

import (
	"testing"
)

// BenchmarkConvertListToMap1-12           20000000                99.7 ns/op             0 B/op          0 allocs/op
// BenchmarkConvertListToMap2-12           20000000                98.2 ns/op             0 B/op          0 allocs/op
// BenchmarkConvertListToMap3-12           20000000               108 ns/op               0 B/op          0 allocs/op
// BenchmarkConvertListToMap4-12           20000000                97.7 ns/op             0 B/op          0 allocs/op
// BenchmarkConvertListToMap5-12            5000000               311 ns/op              80 B/op          5 allocs/op

type Hoge struct {
	Name string
}

var users = []Hoge{
	Hoge{Name: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHoge1"},
	Hoge{Name: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHoge2"},
	Hoge{Name: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHoge3"},
	Hoge{Name: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHoge3"},
	Hoge{Name: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHoge3"},
}
var pusers = []*Hoge{
	&Hoge{Name: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHoge1"},
	&Hoge{Name: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHoge2"},
	&Hoge{Name: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHoge3"},
	&Hoge{Name: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHoge3"},
	&Hoge{Name: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHoge3"},
}

func BenchmarkConvertListToMap1(b *testing.B) {
	userMap := map[string]*Hoge{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j, user := range users {
			userMap[user.Name] = &users[j]
		}
	}
}

func BenchmarkConvertListToMap2(b *testing.B) {
	userMap := map[string]*Hoge{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, user := range pusers {
			userMap[user.Name] = user
		}
	}
}

func BenchmarkConvertListToMap3(b *testing.B) {
	userMap := map[string]*Hoge{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		len := len(pusers)
		for j := 0; j < len; j++ {
			userMap[pusers[j].Name] = pusers[j]
		}
	}
}

func BenchmarkConvertListToMap4(b *testing.B) {
	userMap := map[string]*Hoge{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		len := len(users)
		for j := 0; j < len; j++ {
			userMap[users[j].Name] = &users[j]
		}
	}
}

func BenchmarkConvertListToMap5(b *testing.B) {
	userMap := map[string]*Hoge{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range users {
			// This is slow
			user := users[j]
			userMap[user.Name] = &user
		}
	}
}

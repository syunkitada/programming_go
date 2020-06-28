package main

import "testing"

// むやみにmapを使う前に、まずはarrayやstructで代用できないか検討するとよい
// stringの探索は遅いので、できるだけintで代用できないか検討するとよい

// BenchmarkArray-12               2000000000               0.27 ns/op            0 B/op          0 allocs/op
// BenchmarkCounterStruct-12       2000000000               0.26 ns/op            0 B/op          0 allocs/op
// BenchmarkMap-12                 20000000                78.8 ns/op             0 B/op          0 allocs/op
// BenchmarkIntMap-12              20000000                69.1 ns/op             0 B/op          0 allocs/op
// BenchmarkReadArray-12           2000000000               0.27 ns/op            0 B/op          0 allocs/op
// BenchmarkReadMap-12             20000000                59.2 ns/op             0 B/op          0 allocs/op
// BenchmarkScanArray-12           20000000                64.7 ns/op             0 B/op          0 allocs/op
// BenchmarkScanIntArray-12        100000000               15.3 ns/op             0 B/op          0 allocs/op

const (
	Counter0 = 0
	Counter1 = 1
	Counter2 = 2
	Counter3 = 3
	Counter4 = 4
)

type CounterData struct {
	Counter0 int
	Counter1 int
	Counter2 int
	Counter3 int
	Counter4 int
}

func BenchmarkArray(b *testing.B) {
	b.ResetTimer()
	var counter [5]int
	for i := 0; i < b.N; i++ {
		counter[Counter0] = 1
		counter[Counter1] = 1
		counter[Counter2] = 1
		counter[Counter3] = 1
		counter[Counter4] = 1
	}
}

func BenchmarkCounterStruct(b *testing.B) {
	b.ResetTimer()
	data := CounterData{}
	for i := 0; i < b.N; i++ {
		data.Counter0 = 1
		data.Counter1 = 1
		data.Counter2 = 1
		data.Counter3 = 1
		data.Counter4 = 1
	}
}

func BenchmarkMap(b *testing.B) {
	b.ResetTimer()
	counter := map[string]int{}
	for i := 0; i < b.N; i++ {
		counter["Counter0"] = 1
		counter["Counter1"] = 1
		counter["Counter2"] = 1
		counter["Counter3"] = 1
		counter["Counter4"] = 1
	}
}

func BenchmarkIntMap(b *testing.B) {
	b.ResetTimer()
	counter := map[int]int{}
	for i := 0; i < b.N; i++ {
		counter[Counter0] = 1
		counter[Counter1] = 1
		counter[Counter2] = 1
		counter[Counter3] = 1
		counter[Counter4] = 1
	}
}

func BenchmarkReadArray(b *testing.B) {
	var counter [5]int
	counter[Counter0] = 1
	counter[Counter1] = 1
	counter[Counter2] = 1
	counter[Counter3] = 1
	counter[Counter4] = 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = counter[Counter0]
		_ = counter[Counter1]
		_ = counter[Counter2]
		_ = counter[Counter3]
		_ = counter[Counter4]
	}
}

func BenchmarkReadMap(b *testing.B) {
	counter := map[string]int{}
	counter["Counter0"] = 1
	counter["Counter1"] = 1
	counter["Counter2"] = 1
	counter["Counter3"] = 1
	counter["Counter4"] = 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = counter["Counter0"]
		_ = counter["Counter1"]
		_ = counter["Counter2"]
		_ = counter["Counter3"]
		_ = counter["Counter4"]
	}
}

type Counter struct {
	Name    string
	IntName int
	Counter int
}

func BenchmarkScanArray(b *testing.B) {
	counters := []Counter{
		Counter{Name: "Counter0", Counter: 1},
		Counter{Name: "Counter1", Counter: 1},
		Counter{Name: "Counter2", Counter: 1},
		Counter{Name: "Counter3", Counter: 1},
		Counter{Name: "Counter4", Counter: 1},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetCounter(counters, "Counter0")
		_ = GetCounter(counters, "Counter1")
		_ = GetCounter(counters, "Counter2")
		_ = GetCounter(counters, "Counter3")
		_ = GetCounter(counters, "Counter4")
	}
}

func GetCounter(counters []Counter, name string) *Counter {
	for i, counter := range counters {
		if counter.Name == name {
			return &counters[i]
		}
	}
	return nil
}

func BenchmarkScanIntArray(b *testing.B) {
	counters := []Counter{
		Counter{IntName: Counter0, Counter: 1},
		Counter{IntName: Counter1, Counter: 1},
		Counter{IntName: Counter2, Counter: 1},
		Counter{IntName: Counter3, Counter: 1},
		Counter{IntName: Counter4, Counter: 1},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetCounterByInt(counters, Counter0)
		_ = GetCounterByInt(counters, Counter1)
		_ = GetCounterByInt(counters, Counter2)
		_ = GetCounterByInt(counters, Counter3)
		_ = GetCounterByInt(counters, Counter4)
	}
}

func GetCounterByInt(counters []Counter, intName int) *Counter {
	for i, counter := range counters {
		if counter.IntName == intName {
			return &counters[i]
		}
	}
	return nil
}

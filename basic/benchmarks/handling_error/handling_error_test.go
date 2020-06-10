package handling_error

import (
	"fmt"
	"testing"
)

// BenchmarkHandlingError1-12      50000000                31.9 ns/op             0 B/op          0 allocs/op
// BenchmarkHandlingError2-12      20000000                78.0 ns/op            16 B/op          1 allocs/op

func ReturnErrorInternal() error {
	return fmt.Errorf("")
}

func ReturnError1() (err error) {
	err = ReturnErrorInternal()
	return
}

func ReturnError2() error {
	if err := ReturnErrorInternal(); err != nil {
		return err
	}
	return nil
}

func BenchmarkHandlingError1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ReturnError1()
	}
}

func BenchmarkHandlingError2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ReturnError2()
	}
}

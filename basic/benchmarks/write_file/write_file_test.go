package write_file

import (
	"io/ioutil"
	"os"
	"testing"
)

// BenchmarkWriteFile1-12             50000             30785 ns/op             128 B/op          4 allocs/op
// BenchmarkWriteFile2-12             50000             31022 ns/op             128 B/op          4 allocs/op

const tmpFile = "/tmp/test"

func BenchmarkWriteFile1(b *testing.B) {
	tmp := []byte("Hello World")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		file, err := os.Create(tmpFile)
		if err != nil {
			return
		}

		_, err = file.Write(tmp)
		if err != nil {
			return
		}
		file.Close()
	}
}

func BenchmarkWriteFile2(b *testing.B) {
	tmp := []byte("Hello World")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := ioutil.WriteFile(tmpFile, tmp, 0644)
		if err != nil {
			return
		}
	}
}

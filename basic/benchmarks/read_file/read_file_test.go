package read_file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// BenchmarkReadFileAll1-12          100000             13253 ns/op            4216 B/op          4 allocs/op
// BenchmarkReadFileAll2-12          100000             16992 ns/op            5848 B/op         55 allocs/op
// BenchmarkReadFileAll3-12          100000             20781 ns/op            8904 B/op          9 allocs/op

const meminfoFile = "/proc/meminfo"

func BenchmarkReadFileAll1(b *testing.B) {
	var tmp []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		file, err := os.Open(meminfoFile)
		if err != nil {
			return
		}

		reader := bufio.NewReader(file)
		for {
			tmp, _, err = reader.ReadLine()
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
		}
		file.Close()
	}
	_ = tmp
}

func BenchmarkReadFileAll2(b *testing.B) {
	var tmp string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		file, err := os.Open(meminfoFile)
		if err != nil {
			return
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			tmp = scanner.Text()
		}
		file.Close()
	}
	_ = tmp
}

func BenchmarkReadFileAll3(b *testing.B) {
	var tmp string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes, err := ioutil.ReadFile(meminfoFile)
		if err != nil {
			return
		}
		for _, str := range strings.Split(string(bytes), "\n") {
			tmp = str
		}
	}
	_ = tmp
}

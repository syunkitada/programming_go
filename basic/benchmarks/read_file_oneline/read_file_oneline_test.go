package read_file

import (
	"bufio"
	"io/ioutil"
	"os"
	"testing"
)

// BenchmarkReadFileOneline1-12              200000              7732 ns/op            4216 B/op          4 allocs/op
// BenchmarkReadFileOneline2-12              200000              8278 ns/op            2376 B/op          6 allocs/op
// BenchmarkReadFileOneline3-12              200000              8631 ns/op            4232 B/op          5 allocs/op

const uptimeFile = "/proc/uptime"

// $ cat /proc/uptime
// 3764.69 45088.40

func BenchmarkReadFileOneline1(b *testing.B) {
	var err error
	var file *os.File
	var tmp []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		file, err = os.Open(uptimeFile)
		if err != nil {
			return
		}

		reader := bufio.NewReader(file)
		tmp, _, err = reader.ReadLine()
		file.Close()
	}
	_ = tmp
}

func BenchmarkReadFileOneline2(b *testing.B) {
	var err error
	var tmp []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmp, err = ioutil.ReadFile(uptimeFile)
		if err != nil {
			return
		}
	}
	_ = tmp
}

func BenchmarkReadFileOneline3(b *testing.B) {
	var err error
	var file *os.File
	var tmp string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		file, err = os.Open(uptimeFile)
		if err != nil {
			return
		}

		scanner := bufio.NewScanner(file)
		scanner.Scan()
		tmp = scanner.Text()
		file.Close()
	}
	_ = tmp
}

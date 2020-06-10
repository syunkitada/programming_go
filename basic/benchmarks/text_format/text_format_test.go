package text_format

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"gopkg.in/yaml.v2"
)

// BenchmarkJson1-12        1000000              1475 ns/op             208 B/op          4 allocs/op
// BenchmarkJson2-12        1000000              1490 ns/op             208 B/op          4 allocs/op
// BenchmarkYaml1-12         100000             22183 ns/op            9968 B/op         69 allocs/op
// BenchmarkYaml2-12         100000             20649 ns/op            9968 B/op         69 allocs/op
// BenchmarkXml1-12          100000             12174 ns/op            6024 B/op         41 allocs/op
// BenchmarkXml2-12          100000             12266 ns/op            6024 B/op         41 allocs/op

type User1 struct {
	Name string
	Age  int
}

type User2 struct {
	Name string `json:"name" yaml:"name" xml:"name"`
	Age  int    `json:"age" yaml:"age" xml:"age"`
}

func BenchmarkJson1(b *testing.B) {
	user := User1{Name: "Hoge", Age: 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes, _ := json.Marshal(&user)
		_ = json.Unmarshal(bytes, &user)
	}
}

func BenchmarkJson2(b *testing.B) {
	user := User2{Name: "Hoge", Age: 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes, _ := json.Marshal(&user)
		_ = json.Unmarshal(bytes, &user)
	}
}

func BenchmarkYaml1(b *testing.B) {
	user := User1{Name: "Hoge", Age: 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes, _ := yaml.Marshal(&user)
		_ = yaml.Unmarshal(bytes, &user)
	}
}

func BenchmarkYaml2(b *testing.B) {
	user := User2{Name: "Hoge", Age: 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes, _ := yaml.Marshal(&user)
		_ = yaml.Unmarshal(bytes, &user)
	}
}

func BenchmarkXml1(b *testing.B) {
	user := User1{Name: "Hoge", Age: 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes, _ := xml.Marshal(&user)
		_ = xml.Unmarshal(bytes, &user)
	}
}

func BenchmarkXml2(b *testing.B) {
	user := User2{Name: "Hoge", Age: 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes, _ := xml.Marshal(&user)
		_ = xml.Unmarshal(bytes, &user)
	}
}

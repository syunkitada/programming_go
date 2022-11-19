package main

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Data struct {
	Name string
}

func BenchmarkEncode(b *testing.B) {
	data := Data{
		Name: "hoge",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(&data)
	}
}

func BenchmarkDecode(b *testing.B) {
	data := Data{}
	bytes := []byte(`{"Name": "hoge"}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bytes, &data)
	}
}

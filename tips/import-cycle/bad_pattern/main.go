package main

import "github.com/syunkitada/go-samples/tips/import-cycle/bad_pattern/piyo"

func main() {
	piyo := piyo.New()
	piyo.Hoge()
}

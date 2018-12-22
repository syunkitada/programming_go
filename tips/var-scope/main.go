package main

import (
	"fmt"
)

func main() {
	var err error
	defer fmt.Printf("defer %v\n", err)
	defer func() {
		defer fmt.Printf("defer func %v\n", err)
	}()

	hoge, err := hogeErr()
	fmt.Println(hoge, err)

	func() {
		piyo, err := piyoErr()
		fmt.Println(piyo, err)
	}()
}

func hogeErr() (int, error) {
	return 0, fmt.Errorf("hoge")
}

func piyoErr() (int, error) {
	return 1, fmt.Errorf("piyo")
}

package main

import (
	"errors"
)

func BenchmarkArrocateError1() {
	var err error
	for i := 0; i < 100; i++ {
		if err = errors.New("TEST"); err != nil {
			continue
		}
	}
}

func BenchmarkArrocateError2() {
	for i := 0; i < 100; i++ {
		if err := errors.New("TEST"); err != nil {
			continue
		}
	}
}

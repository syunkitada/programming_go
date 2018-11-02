package main

import (
	"fmt"
	"unsafe"
)

type Sample struct {
	flag      bool
	num       int
	num32     int32
	num64     int64
	str       string
	numArray5 [5]int
	strArray5 [5]string
	numSlice  []int
	strSlice  []string
	numMap    map[string]int
}

func main() {
	sample := Sample{}
	fmt.Printf("Sample: %v\n", unsafe.Sizeof(sample))
	fmt.Printf("*Sample: %v\n", unsafe.Sizeof(&sample))                   // 8
	fmt.Printf("Sample.flag: %v\n", unsafe.Sizeof(sample.flag))           // 1
	fmt.Printf("Sample.num: %v\n", unsafe.Sizeof(sample.num))             // 8
	fmt.Printf("Sample.num32: %v\n", unsafe.Sizeof(sample.num32))         // 4
	fmt.Printf("Sample.num64: %v\n", unsafe.Sizeof(sample.num64))         // 8
	fmt.Printf("Sample.str: %v\n", unsafe.Sizeof(sample.str))             // 16
	fmt.Printf("Sample.numArray5: %v\n", unsafe.Sizeof(sample.numArray5)) // 40
	fmt.Printf("Sample.strArray5: %v\n", unsafe.Sizeof(sample.strArray5)) // 80
	fmt.Printf("Sample.numSlice: %v\n", unsafe.Sizeof(sample.numSlice))   // 24
	fmt.Printf("Sample.strSlice: %v\n", unsafe.Sizeof(sample.strSlice))   // 24
	fmt.Printf("Sample.numMap: %v\n", unsafe.Sizeof(sample.numMap))       // 8

	for i := 1; i < 100; i += 1 {
		sample.numSlice = append(sample.numSlice, i)
		sample.strSlice = append(sample.strSlice, "a")
		sample.str += "a"
	}
	fmt.Println("------------------------------------------------------")
	fmt.Printf("Sample.numSlice: %v\n", unsafe.Sizeof(sample.numSlice)) // 24
	fmt.Printf("Sample.strSlice: %v\n", unsafe.Sizeof(sample.strSlice)) // 24
	fmt.Printf("Sample.str: %v\n", unsafe.Sizeof(sample.str))           // 16
}

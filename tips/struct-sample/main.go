package main

import (
	"fmt"
)

type Base struct {
	num int
}

func (b *Base) set(num int) {
	b.num = num
}

func (b *Base) get() int {
	return b.num
}

type Sub struct {
	Base
}

func main() {
	base := Base{}
	fmt.Println(base.get())

	sub := Sub{
		Base: Base{
			num: 1,
		},
	}
	fmt.Println(sub.get())
	sub.set(2)
	fmt.Println(sub.get())
}

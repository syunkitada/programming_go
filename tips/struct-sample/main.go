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

func (b *Base) dump() {
	fmt.Printf("Base %v\n", b.num)
}

func (b *Base) callDump() {
	b.dump()
}

type Sub struct {
	Base
}

func (s Sub) dump() {
	fmt.Printf("Sub %v\n", s.num)
}

func main() {
	base := Base{}
	base.dump()
	base.callDump()

	sub := Sub{
		Base: Base{
			num: 1,
		},
	}
	sub.dump()

	sub.set(2)
	sub.dump()
	sub.callDump()
}

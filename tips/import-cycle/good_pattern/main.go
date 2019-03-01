package main

import (
	"github.com/syunkitada/go-samples/tips/import-cycle/good_pattern/hoge"
	"github.com/syunkitada/go-samples/tips/import-cycle/good_pattern/piyo"
)

func main() {
	var h *hoge.Hoge
	var p *piyo.Piyo
	h = hoge.New(p)
	p = piyo.New(h)

	p.Hoge()
	h.Piyo()
}

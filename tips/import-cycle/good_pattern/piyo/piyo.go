package piyo

import (
	"fmt"

	"github.com/syunkitada/go-samples/tips/import-cycle/good_pattern/hoge_interface"
)

type Piyo struct {
	hoge hoge_interface.HogeInterface
}

func New(hoge hoge_interface.HogeInterface) *Piyo {
	return &Piyo{
		hoge: hoge,
	}
}

func (piyo *Piyo) Piyo() {
	fmt.Println("piyo")
}

func (piyo *Piyo) Hoge() {
	piyo.hoge.Hoge()
}

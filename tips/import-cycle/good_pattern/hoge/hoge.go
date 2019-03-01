package hoge

import (
	"fmt"

	"github.com/syunkitada/go-samples/tips/import-cycle/good_pattern/piyo_interface"
)

type Hoge struct {
	piyo piyo_interface.PiyoInterface
}

func New(piyo piyo_interface.PiyoInterface) *Hoge {
	return &Hoge{
		piyo: piyo,
	}
}

func (hoge *Hoge) Hoge() {
	fmt.Println("hoge")
}

func (hoge *Hoge) Piyo() {
	hoge.piyo.Piyo()
}

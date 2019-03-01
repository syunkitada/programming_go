package piyo

import (
	"fmt"

	"github.com/syunkitada/go-samples/tips/import-cycle/bad_pattern/hoge"
)

type Piyo struct{}

func New() *Piyo {
	return &Piyo{}
}

func (piyo *Piyo) Piyo() {
	fmt.Println("piyo")
}

func (piyo *Piyo) Hoge() {
	hoge := hoge.New()
	hoge.Hoge()
}

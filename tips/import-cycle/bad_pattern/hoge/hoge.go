package hoge

import (
	"fmt"

	"github.com/syunkitada/go-samples/tips/import-cycle/bad_pattern/piyo"
)

type Hoge struct{}

func New() *Hoge {
	return &Hoge{}
}

func (hoge *Hoge) Hoge() {
	fmt.Println("hoge")
}

func (hoge *Hoge) Piyo() {
	piyo := piyo.New()
	piyo.Piyo()
}

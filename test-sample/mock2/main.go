package main

import (
	"fmt"

	"github.com/syunkitada/go-samples/test-sample/mock2/client"
)

type Worker struct {
	client client.Client
}

func (worker *Worker) exec(input string) (result string, err error) {
	result, err = worker.client.Request(input)
	return
}

func main() {
	worker := Worker{
		client: client.New(),
	}
	result, err := worker.exec("hoge")
	fmt.Println(result, err)
}

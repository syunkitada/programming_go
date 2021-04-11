package main

import (
	"fmt"

	"github.com/syunkitada/go-samples/test-sample/coverapp/api"
)

func main() {
	fmt.Printf("Result=%d\n", api.GetResult(2, 3))
}

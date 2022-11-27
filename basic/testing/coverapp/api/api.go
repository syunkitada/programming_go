package api

import "github.com/syunkitada/go-samples/test-sample/coverapp/lib"

func GetResult(x, y int) int {
	return lib.Add(x, y)
}

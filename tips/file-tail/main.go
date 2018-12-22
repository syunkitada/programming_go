package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

func main() {
	t, err := tail.TailFile("/home/owner/.goapp/logs/goapp-resource-api.log", tail.Config{Follow: true})
	if err != nil {
		fmt.Print(err)
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/syunkitada/go-samples/tips/runner/runner"
)

func main() {
	fmt.Println("DEBUG main")
	var runners []*runner.Runner
	runners = append(runners, runner.New(&runner.Config{
		Timeout:  30,
		Interval: 60,
		Cmd:      "sleep 3600",
	}))

	runners = append(runners, runner.New(&runner.Config{
		Timeout:  30,
		Interval: 60,
		UseShell: true,
		Cmd:      "sh -c 'sleep 3600'",
	}))

	for _, runner := range runners {
		runner.Start()
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Printf("starting shutdown")
	var wg sync.WaitGroup
	for i := range runners {
		wg.Add(1)
		go func(i int) {
			runners[i].Stop()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("end shutdown")
}

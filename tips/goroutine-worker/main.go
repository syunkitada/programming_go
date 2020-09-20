package main

import (
	"fmt"
	"runtime"
	"time"
)

type Result struct {
	Result int
	Error  error
}

type Worker struct {
	Id int
}

func NewWorker(id int) *Worker {
	return &Worker{Id: id}
}

func (w *Worker) Start(jobs <-chan int, results chan<- *Result) {
	for j := range jobs {
		fmt.Println("worker", w.Id, "processing job", j)
		time.Sleep(1 * time.Second)
		results <- &Result{Result: w.Id}
	}
}

func startWorkers() {
	fmt.Println("Start NumGoroutine ", runtime.NumGoroutine())

	workers := 3
	lenJobs := 10
	jobs := make(chan int, lenJobs)
	results := make(chan *Result, lenJobs)
	for i := 0; i < workers; i++ {
		worker := NewWorker(i)
		go worker.Start(jobs, results)
	}

	fmt.Println("Started NumGoroutine ", runtime.NumGoroutine())

	for i := 0; i < lenJobs; i++ {
		jobs <- i
	}
	// jobsをcloseしないとStartのgoroutineが終了しないので注意
	close(jobs)

	for i := 0; i < lenJobs; i++ {
		result := <-results
		if result.Error != nil {
			fmt.Println("ERROR", result.Error.Error())
		}
	}
	// 厳密にはresultsはcloseしなくてもよい
	// 参照がなくなればGCによって回収されるが明示的に閉じれるなら閉じておくのがよい
	close(results)

	// jobsが消化し終わればworkerのgoroutineも終了する
	fmt.Println("End NumGoroutine ", runtime.NumGoroutine())
}

func main() {
	startWorkers()
}

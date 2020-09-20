package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Task struct {
	Error error
	f     func() error
}

func NewTask(f func() error) *Task {
	return &Task{f: f}
}

func (t *Task) Run(wg *sync.WaitGroup) {
	t.Error = t.f()
	wg.Done()
}

type Pool struct {
	Tasks       []*Task
	concurrency int
	tasksChan   chan *Task
	wg          sync.WaitGroup
}

func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
	}
}

func (p *Pool) RunTasks() {
	for i := 0; i < p.concurrency; i++ {
		go p.StartWorker()
	}
	fmt.Println("Started NumGoroutine ", runtime.NumGoroutine())

	p.wg.Add(len(p.Tasks))
	for _, task := range p.Tasks {
		p.tasksChan <- task
	}

	close(p.tasksChan)

	p.wg.Wait()
}

func (p *Pool) StartWorker() {
	for task := range p.tasksChan {
		task.Run(&p.wg)
	}
}

func main() {
	fmt.Println("Start NumGoroutine ", runtime.NumGoroutine())

	tasks := []*Task{}
	tasks = append(
		tasks,
		NewTask(func() error {
			time.Sleep(2 * time.Second)
			fmt.Println("task1")
			return nil
		}),
	)

	tasks = append(
		tasks,
		NewTask(func() error {
			time.Sleep(1 * time.Second)
			fmt.Println("task2")
			return nil
		}),
	)

	tasks = append(
		tasks,
		NewTask(func() error {
			time.Sleep(2 * time.Second)
			fmt.Println("task3")
			return fmt.Errorf("Failed task3")
		}),
	)

	tasks = append(
		tasks,
		NewTask(func() error {
			fmt.Println("task4")
			return nil
		}),
	)

	concurrency := 2
	p := NewPool(tasks, concurrency)
	p.RunTasks()

	for _, task := range p.Tasks {
		if task.Error != nil {
			fmt.Println("ERROR", task.Error.Error())
		}
	}

	fmt.Println("End NumGoroutine ", runtime.NumGoroutine())
}

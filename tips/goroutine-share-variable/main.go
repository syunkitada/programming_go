package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Name string
}

type Result struct {
	Result string
}

type Worker struct {
	taskMutex sync.Mutex
	tasks     []Task
}

func NewWorker() *Worker {
	return &Worker{}
}

func (w *Worker) startSub() {
	fmt.Println("StartSub")
	time.Sleep(1 * time.Second)
	w.taskMutex.Lock()
	tasks := w.tasks
	w.taskMutex.Unlock()
	fmt.Println("Sub1", len(tasks))
	time.Sleep(3 * time.Second)
	fmt.Println("Sub2", len(tasks))

	fmt.Println("EndSub")
}

func (w *Worker) startMain() {
	fmt.Println("StartMain")
	w.taskMutex.Lock()
	w.tasks = []Task{
		Task{Name: "task1"},
	}
	w.taskMutex.Unlock()
	fmt.Println("Main1", len(w.tasks))
	time.Sleep(2 * time.Second)

	w.taskMutex.Lock()
	w.tasks = []Task{
		Task{Name: "task1"},
		Task{Name: "task2"},
	}
	w.taskMutex.Unlock()
	fmt.Println("EndMain")
}

func main() {
	worker := NewWorker()
	go worker.startSub()
	worker.startMain()
	time.Sleep(10 * time.Second)
}

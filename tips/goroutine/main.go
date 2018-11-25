package main

import (
	"log"
	"runtime"
	"sync"
	"time"

	"golang.org/x/net/context"
)

func singleGoroutine() {
	// 一つのgoroutineを管理する場合、その実行の結果を得るためにchannelを利用する

	// channelは goroutine 間でのメッセージパッシングをするためのもの
	msg := make(chan string) // メッセージの型を指定できる
	go func() {
		// channelは引数や戻り値に使える
		msg <- "done"
	}()
	result := <-msg // channelは、送受信が完了するまでブロックする(goroutineの終了を待ち、その戻り値を得ることができる
	log.Printf("chanelSample: %v", result)
}

func multiGoroutine() {
	// 複数のgoroutineを管理する場合、WaitGroupを利用する
	wg := sync.WaitGroup{}

	// wg.Add(1)で、goroutineを生成するたびインクリメントし、
	// 各goroutine内でwg.Doneで、デクリメントすることで、すべてのgoroutineが終了したかを判定している

	wg.Add(1)
	go echoSleep(&wg, "hoge", 10)

	wg.Add(1)
	go echoSleep(&wg, "piyo", 5)

	wg.Add(1)
	go echoSleepWithTimeout(&wg, "foo", 5, 10)

	wg.Add(1)
	go echoSleepWithTimeout(&wg, "bar", 10, 5)

	time.Sleep(1 * time.Second) // goroutineの起動に若干時間かかるため少し待つ
	log.Printf("Debug: NumGoroutine: %v", runtime.NumGoroutine())

	wg.Wait() // ブロックし、全goroutineが終わったら次に進む

	log.Printf("Debug: NumGoroutine: %v", runtime.NumGoroutine())
}

func echoSleep(wg *sync.WaitGroup, name string, sleepDuration int) {
	defer func() { wg.Done() }()

	log.Printf("%v: starting, and sleep %v", name, sleepDuration)
	time.Sleep(time.Duration(sleepDuration) * time.Second)
	log.Printf("%v: done", name)
}

func echoSleepWithTimeout(wg *sync.WaitGroup, name string, sleepDuration int, timeoutDuration int) {
	defer func() { wg.Done() }()

	log.Printf("%v: starting with timeout %v", name, timeoutDuration)
	result := make(chan string)

	// WithTimeoutにより指定時間後にキャンセルを行う
	// 親がキャンセルされると、子もキャンセルされる
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeoutDuration)*time.Second)
	defer cancel()
	go func() {
		log.Printf("%v: starting task, and sleep %v", name, sleepDuration)
		time.Sleep(time.Duration(sleepDuration) * time.Second)
		log.Printf("%v: done", name)
		result <- "Success"
	}()

	// 複数のchanelを扱う場合はselectを利用する
	select {
	case r := <-result:
		log.Printf("%v: recieved result: %v", name, r)
	case <-ctx.Done():
		log.Printf("%v: failed: %v", name, ctx.Err())
	}

	log.Printf("%v: done with timeout", name)
}

func main() {
	singleGoroutine()
	multiGoroutine()
}

package goroutine_share_memory

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// BenchmarkShareMemoryWithChannel1-12       500000              2430 ns/op             421 B/op          8 allocs/op
// BenchmarkShareMemoryWithChannel2-12       500000              2580 ns/op             437 B/op          9 allocs/op

// channelはクローズすべきか？
// 使い終わったchannelは、ガベージコレクトされるので明示的に閉じる必要はない
// https://stackoverflow.com/questions/8593645/is-it-ok-to-leave-a-channel-open

func BenchmarkShareMemoryWithChannel1(b *testing.B) {
	var tmp []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ShareMemoryWithChannel1()
	}
	_ = tmp
}

func ShareMemoryWithChannel1() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	errChan := make(chan error)
	go func() {
		errChan <- fmt.Errorf("Error")
	}()

	select {
	case err = <-errChan:
		return
	case <-ctx.Done():
		err = fmt.Errorf("Timeout")
		return
	}
	return
}

func BenchmarkShareMemoryWithChannel2(b *testing.B) {
	var tmp []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ShareMemoryWithChannel2()
	}
	_ = tmp
}

// channelをちゃんと閉じるパターン
func ShareMemoryWithChannel2() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tmpChan := make(chan int)
	go func() {
		err = fmt.Errorf("Error")
		close(tmpChan)
	}()

	select {
	case <-tmpChan:
		return
	case <-ctx.Done():
		err = fmt.Errorf("Timeout")
		return
	}
	return
}

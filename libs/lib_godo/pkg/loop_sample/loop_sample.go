package loop_sample

import (
	"fmt"
	"time"
)

func Main() {
	for {
		fmt.Printf("%v: hello\n", time.Now())
		time.Sleep(2 * time.Second)
	}
}

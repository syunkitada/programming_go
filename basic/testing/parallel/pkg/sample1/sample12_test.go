package sample11

import (
	"fmt"
	"testing"
	"time"
)

func TestSample12(t *testing.T) {
	fmt.Println("TestSample12", time.Now())
	time.Sleep(1 * time.Second)
	actual := sample12()
	expected := "sample12"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
	fmt.Println("TestSample12 End", time.Now())
}

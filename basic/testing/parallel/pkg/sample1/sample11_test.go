package sample11

import (
	"fmt"
	"testing"
	"time"
)

func TestSample11(t *testing.T) {
	fmt.Println("TestSample11", time.Now())
	time.Sleep(1 * time.Second)

	actual := sample11()
	expected := "sample11"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
	fmt.Println("TestSample11 End", time.Now())
}

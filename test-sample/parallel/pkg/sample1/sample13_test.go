package sample11

import (
	"fmt"
	"testing"
	"time"
)

func TestSample13(t *testing.T) {
	fmt.Println("TestSample13", time.Now())
	time.Sleep(1 * time.Second)
	actual := sample13()
	expected := "sample13"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
	fmt.Println("TestSample13 End", time.Now())
}

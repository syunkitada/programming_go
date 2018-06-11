package simple_test

import (
	"github.com/syunkitada/go-sample/pkg/test_example/simple"
	"testing"
)

func TestHello(t *testing.T) {
	actual := simple.Hello()
	expected := "hello"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

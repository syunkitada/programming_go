package main

import (
	"fmt"
	"testing"

	"github.com/syunkitada/go-samples/test-sample/mock2/client"
)

func TestWorker(t *testing.T) {
	mockClientPatch := client.MockClientPatch{
		Request: func(input string) (response string, err error) {
			response = fmt.Sprintf("hello %s", input)
			return
		},
	}
	mockClient := client.NewMockClient(&mockClientPatch)

	worker := Worker{client: mockClient}

	result, err := worker.exec("hoge")
	if err != nil {
		t.Errorf("error occurerd: %s", err.Error())
	}
	if result != "hello hoge" {
		t.Errorf("unexpected result: %s", result)
	}

	// 途中で動作を変えたい場合
	mockClientPatch.Request = func(input string) (response string, err error) {
		response = fmt.Sprintf("goodby %s", input)
		return
	}
	result, err = worker.exec("hoge")
	if err != nil {
		t.Errorf("error occurerd: %s", err.Error())
	}
	if result != "goodby hoge" {
		t.Errorf("unexpected result: %s", result)
	}
}

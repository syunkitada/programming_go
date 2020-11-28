package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/syunkitada/go-samples/test-sample/mock1/client"
)

func TestWorker(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockClient := client.NewMockClient(ctrl)

	worker := Worker{client: mockClient}

	mockClient.
		EXPECT().
		Request(gomock.Eq("hoge")).
		Return("hello, hoge", nil)

	result, err := worker.exec("hoge")
	if err != nil {
		t.Errorf("error occurerd: %s", err.Error())
	}
	if result != "hello, hoge" {
		t.Errorf("unexpected result: %s", result)
	}
}

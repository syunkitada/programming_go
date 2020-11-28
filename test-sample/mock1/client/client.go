//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package client

import "fmt"

type Client interface {
	Request(input string) (response string, err error)
}

type ApiClient struct {
}

func New() Client {
	return &ApiClient{}
}

func (ApiClient) Request(input string) (response string, err error) {
	response = fmt.Sprintf("hello %s", input)
	return
}

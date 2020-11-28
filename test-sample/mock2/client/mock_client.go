package client

type MockClient struct {
	patch *MockClientPatch
}

type MockClientPatch struct {
	Request func(input string) (response string, err error)
}

func NewMockClient(patch *MockClientPatch) *MockClient {
	return &MockClient{patch: patch}
}

func (client *MockClient) Request(input string) (response string, err error) {
	return client.patch.Request(input)
}

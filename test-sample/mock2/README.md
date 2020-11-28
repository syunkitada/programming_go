# mock1

- mockgen を利用せず、自前で mock を作る場合
- mock に状態を持たせるなど柔軟なことをやる場合には自前で mock を作ったほうがよい
- mock 側は以下のように patch をポインタで持たせて test 側で自由に挙動を変えられるようにしておく

```
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

```

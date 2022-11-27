# sample1

- [公式ドキュメントの Tutorial](https://goswagger.io/tutorial/todo-list.html)
- sagger.yml からコード生成するパターン

```
# validate
$ swagger validate ./swagger.yml
```

```
# ymlを修正したら以下でコードを自動生成します
$ swagger generate server -A todo-list -f ./swagger.yml
$ swagger generate client -A todo-list -f ./swagger.yml
$ swagger generate cli -A todo-list -f ./swagger.yml

# コードの依存をインストールしておく
$ go mod tidy

# 実行
$ go run cmd/todo-list-server/main.go -h
```

```
$ tree
.
├── README.md
├── cmd
│   └── todo-list-server
│       └── main.go
├── models
│   ├── error.go
│   └── item.go
├── restapi
│   ├── configure_todo_list.go
│   ├── doc.go
│   ├── embedded_spec.go
│   ├── operations
│   │   ├── todo_list_api.go
│   │   └── todos
│   │       ├── get.go
│   │       ├── get_parameters.go
│   │       ├── get_responses.go
│   │       └── get_urlbuilder.go
│   └── server.go
└── swagger.yml
```

```
$ go run cmd/todo-list-server/main.go --port 1080
2022/10/10 13:45:07 Serving todo list at http://127.0.0.1:1080

$ curl localhost:1080/
"operation todos.Get has not yet been implemented"
```

```
$ curl -XPOST localhost:1080 -d '{"description":"message hoge"}' -H 'content-type: application/json'
{"code":500,"message":"no consumer registered for application/json"}

$ curl -XPOST localhost:1080 -d '{"description":"message hoge"}' -H 'content-type: application/io.goswagger.examples.todo-list.v1+json'
{"description":"message hoge","id":3}


$ curl localhost:1080
[{"description":"message hoge","id":1},{"description":"message hoge","id":2},{"description":"message hoge","id":3}]

$ curl -XPUT localhost:1080/2/ -d '{"description":"message hogepiyo"}' -H 'content-type: application/io.goswagger.examples.todo-list.v1+json'
{"description":"message hogepiyo","id":2}

$ curl localhost:1080
[{"description":"message hoge","id":1},{"description":"message hogepiyo","id":2},{"description":"message hoge","id":3}]

$ curl -XDELETE localhost:1080/3/

$ curl localhost:1080
[{"description":"message hoge","id":1},{"description":"message hogepiyo","id":2}]
```

```
$ go run cmd/cli/main.go --hostname localhost:1080 todos addOne --item.description hoge
{"description":"hoge","id":1}

$ go run cmd/cli/main.go --hostname localhost:1080 todos findTodos
[{"description":"hoge","id":1},{"description":"piyo","id":2}]
```

## コードから spec を生成するパターン

```
$ swagger generate spec -o ./swagger.json
```

- spec はコードのコメントアウトのタグを頼りに自動生成されます

```
$ grep 'swagger:' * -r | grep .go
models/error.go:// swagger:model error
models/item.go:// swagger:model item
restapi/doc.go:// swagger:meta
restapi/operations/todos/get.go:        Get swagger:route GET / todos get
restapi/operations/todos/add_one_responses.go:swagger:response addOneCreated
restapi/operations/todos/add_one_responses.go:swagger:response addOneDefault
restapi/operations/todos/update_one_responses.go:swagger:response updateOneOK
restapi/operations/todos/update_one_responses.go:swagger:response updateOneDefault
restapi/operations/todos/add_one.go:    AddOne swagger:route POST / todos addOne
restapi/operations/todos/get_responses.go:swagger:response getOK
restapi/operations/todos/get_responses.go:swagger:response getDefault
restapi/operations/todos/add_one_parameters.go:// swagger:parameters addOne
restapi/operations/todos/destroy_one.go:        DestroyOne swagger:route DELETE /{id} todos destroyOne
restapi/operations/todos/destroy_one_responses.go:swagger:response destroyOneNoContent
restapi/operations/todos/destroy_one_responses.go:swagger:response destroyOneDefault
restapi/operations/todos/get_parameters.go:// swagger:parameters Get
restapi/operations/todos/find_todos_parameters.go:// swagger:parameters findTodos
restapi/operations/todos/find_todos.go: FindTodos swagger:route GET / todos findTodos
restapi/operations/todos/update_one_parameters.go:// swagger:parameters updateOne
restapi/operations/todos/find_todos_responses.go:swagger:response findTodosOK
restapi/operations/todos/find_todos_responses.go:swagger:response findTodosDefault
restapi/operations/todos/destroy_one_parameters.go:// swagger:parameters destroyOne
restapi/operations/todos/update_one.go: UpdateOne swagger:route PUT /{id} todos updateOne
```

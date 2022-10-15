# sample1

```
# protocol bufferのコード生成
$ make gen
```

# Go Run

```
$ go run cmd/grpc-server/main.go --config-dir=${PWD}/ci/etc
$ go run cmd/grpc-client/main.go --config-dir=${PWD}/ci/etc status
```

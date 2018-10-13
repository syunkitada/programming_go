# grcp-sample


# Go Run
```
$ go run cmd/grpc-server/main.go --config-dir=${PWD}/ci/etc
$ go run cmd/grpc-client/main.go --config-dir=${PWD}/ci/etc status
```


# テスト実行
```
$ go test pkg/ctl/main_test.go
ok      command-line-arguments  0.001s
```


# パッケージ作成
```
$ make rpm
```

# godo-sample


## Setup Development Environment
```
$ dep ensure
```


## Go Run loop-sample with godo
* ファイ変ル更を監視しながら、loop-sampleを実行する
* ファイルが変更されると、自動でloop-sampleが再実行される
```
$ go run cmd/godo-sample/main.go loop-sample --watch
```

# 循環参照

* 循環参照すると以下のようにコンパイルエラーとなるため、インターフェイスを使ってうまく回避する必要がある

```
$ go run bad_pattern/main.go
import cycle not allowed
package main
        imports github.com/syunkitada/go-samples/tips/import-cycle/bad_pattern/piyo
                imports github.com/syunkitada/go-samples/tips/import-cycle/bad_pattern/hoge
                        imports github.com/syunkitada/go-samples/tips/import-cycle/bad_pattern/piyo

```

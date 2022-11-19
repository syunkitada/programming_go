# json

- go-json
  - https://github.com/goccy/go-json
  - 最速？
  - 依存がない
  - json ライブラリと互換なのでライブラリを入れ替えるだけで使える
  - 参考
    - https://engineering.mercari.com/blog/entry/1599563768-081104c850/
    - https://www.youtube.com/watch?v=BpKBUAqkxsw
  - メモ: 高速化の仕組み
    - Zero-allocation
      - zap や zerolog などが使ってるやつ
      - Zero-allocation == Zero-GC
        - GC(STW)の時間は無視できない
        - アロケーション回数 1 回と 0 回では、GC が発生する・しない関わる重要な分かれ目
      - Zero-allocation がなぜ可能か
        - Go コンマイらが賢くオブジェクトの生存期間を解析して、ヒープに確保する必要がないものはスタック確保してくれる仕組み（エスケープ解析）
          -sync.Pool によるオブジェクトの再利用
      - エスケープ解析に優しいコードを書く
- jsoniter
  - https://github.com/json-iterator/go

```
$ go test -bench . ./lib_go-json/... -benchmem
goos: linux
goarch: amd64
pkg: github.com/syunkitada/programming_go/libs/json/standard/lib_go-json
cpu: AMD Ryzen 5 2600 Six-Core Processor
BenchmarkEncode-12       9487821               114.4 ns/op            16 B/op          1 allocs/op
BenchmarkDecode-12       8214225               142.5 ns/op            24 B/op          1 allocs/op
PASS
ok      github.com/syunkitada/programming_go/libs/json/standard/lib_go-json     2.546s
```

```
$ go test -bench . ./lib_jsonitor/... -benchmem
goos: linux
goarch: amd64
pkg: github.com/syunkitada/programming_go/libs/json/standard/lib_jsonitor
cpu: AMD Ryzen 5 2600 Six-Core Processor
BenchmarkEncode-12       5059154               241.7 ns/op            24 B/op          2 allocs/op
BenchmarkDecode-12       5443330               212.3 ns/op             8 B/op          2 allocs/op
PASS
ok      github.com/syunkitada/programming_go/libs/json/standard/lib_jsonitor    2.856s
```

```
$ go test -bench . ./standard/... -benchmem
goos: linux
goarch: amd64
pkg: github.com/syunkitada/programming_go/libs/json/standard/standard
cpu: AMD Ryzen 5 2600 Six-Core Processor
BenchmarkEncode-12       6499034               195.0 ns/op            16 B/op          1 allocs/op
BenchmarkDecode-12       1000000              1059 ns/op             224 B/op          5 allocs/op
PASS
ok      github.com/syunkitada/programming_go/libs/json/standard/standard        2.531s
```

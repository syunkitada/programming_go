# Logger

## Logger 選択時に考慮すること

- 構造化ロギング
  - JSON などのプログラム的に処理しやすいフォーマットを利用できるロギングのこと
  - フォーマットの種類
    - JSON
    - MessagePack
    - Protobuf
    - CBOR
- 高パフォーマンス
  - ログ量によってはその出力自体がアプリケーションに影響があるため、高パフォーマンスを意識して設計されていること
  - Zap では、構造化ロギング時によく利用される interface を利用した reflection ベースのシリアライズや文字列のフォーマットをやめている
    - これより、CPU 消費やメモリアロケーションを大幅に削減できる
  - また、極力ヒープを利用しないように設計されており、メモリ効率がよい
  - Zap の登場以降では後続の Logging ライブラリも Zap の実装を参考にパフォーマンスを意識している傾向がある
- サンプリング
  - 特定条件を満たした場合にログを出力する、制限する機能
  - すべてのログを出力した場合、ログ量が多すぎてそれがパフォーマンスに影響が出てしまう場合に利用する
  - 同じログレベルや同じメッセージログを１秒間に出力できる数を制限するなど
- トレーシング機能
  - 各処理にどれだけ時間がかかったかを計測できるようにする
  - Logger ライブラリ自体がこれを備えてはいないので、独自に実装するか、OpenTracing などの分散トレーシング用のライブラリを利用するとよい
    - 例えば OpenTracing+Zap だと、[go-zap](https://github.com/opentracing-contrib/go-zap)というライブラリがある
  - 独自に実装する場合は、トレーシング用の ID を Context で伝搬させてログに仕込めばよい

## ベンチマーク

```
$ go test -bench . ./benchmark/main_test.go -benchmem
goos: linux
goarch: amd64
BenchmarkBuiltinLog-12                    760764              1549 ns/op              16 B/op          1 allocs/op
BenchmarkZapDevelopmentLog-12             278079              3878 ns/op             328 B/op          7 allocs/op
BenchmarkZapProductionLog-12             8065690               154 ns/op               2 B/op          0 allocs/op
BenchmarkZapLog-12                        285618              3869 ns/op             248 B/op          3 allocs/op
BenchmarkZapSamplingLog-12               6773191               162 ns/op               2 B/op          0 allocs/op
BenchmarkZapLogDisableCaller-12           680856              1760 ns/op               0 B/op          0 allocs/op
BenchmarkZapLogZap-12                     286525              4458 ns/op             376 B/op          4 allocs/op
BenchmarkZapSugerfLog-12                  236835              4808 ns/op             296 B/op          4 allocs/op
BenchmarkZapSugerwLog-12                  239828              5009 ns/op             504 B/op          4 allocs/op
BenchmarkZerolog-12                       963723              1225 ns/op               0 B/op          0 allocs/op
BenchmarkZerologUnix-12                   977259              1227 ns/op               0 B/op          0 allocs/op
PASS
ok      command-line-arguments  16.927s
```

## Logger の候補

- Zap
  - ほぼこれ一択
  - パフォーマンス面、機能面、カスタマイズ性で不満なし
  - Uber で開発されてる安心感
- Zerolog
  - パフォーマンスは Zap と同等だが、カスタマイズ性は Zap のほうがよさそう
  - Zap より機能が絞られてるためか Zap よりも若干性能が良い
  - nested なオブジェクトの出力ができなさそう

## 参考

- [Go ロギングライブラリ](https://qiita.com/nakaryooo/items/2ee140cf4aafa9ff1732)
- [Go のロギングライブラリ 2021 年冬](https://moriyoshi.hatenablog.com/entry/2021/12/14/183703)

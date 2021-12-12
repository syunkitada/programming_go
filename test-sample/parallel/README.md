# テストの並列実行について

- Go のテストはパッケージ単位で並列で行われる
- パッケージ内は直列で行われる

## 実行例

```
$ go clean -testcache
$ go test -v ./pkg/...
=== RUN   TestSample11
TestSample11 2021-12-12 16:24:01.852958035 +0900 JST m=+0.000566960
TestSample11 End 2021-12-12 16:24:02.853491071 +0900 JST m=+1.001100036
--- PASS: TestSample11 (1.00s)
=== RUN   TestSample12
TestSample12 2021-12-12 16:24:02.853646955 +0900 JST m=+1.001255890
TestSample12 End 2021-12-12 16:24:03.853995549 +0900 JST m=+2.001604494
--- PASS: TestSample12 (1.00s)
=== RUN   TestSample13
TestSample13 2021-12-12 16:24:03.854149519 +0900 JST m=+2.001758434
TestSample13 End 2021-12-12 16:24:04.854419434 +0900 JST m=+3.002028409
--- PASS: TestSample13 (1.00s)
PASS
ok      github.com/syunkitada/go-samples/test-sample/parallel/pkg/sample1       3.004s
=== RUN   TestSample11
TestSample11 2021-12-12 16:24:01.840826245 +0900 JST m=+0.000601746
TestSample11 End 2021-12-12 16:24:02.841408424 +0900 JST m=+1.001184015
--- PASS: TestSample11 (1.00s)
=== RUN   TestSample12
TestSample12 2021-12-12 16:24:02.841522178 +0900 JST m=+1.001297669
TestSample12 End 2021-12-12 16:24:03.841839062 +0900 JST m=+2.001614694
--- PASS: TestSample12 (1.00s)
=== RUN   TestSample13
TestSample13 2021-12-12 16:24:03.841971091 +0900 JST m=+2.001746623
TestSample13 End 2021-12-12 16:24:04.842411778 +0900 JST m=+3.002187369
--- PASS: TestSample13 (1.00s)
PASS
ok      github.com/syunkitada/go-samples/test-sample/parallel/pkg/sample2       3.004s
```

## 全テストを直列で実行する例

- すべてを直列化するには、-p 1 オプションを入れるとできる（より正確には並列数の指定を 1 に指定しているだけ）
  - デフォルトの並列数は、CPU 数: runtime.NumCPU()
- 直列化をすると、パッケージ数が増えてきた場合に実行時間が増えるのでなるべく並列でテストを実行できるように実装したほうがよい
  - どうしても直列でテストしないといけないものは、専用のパッケージ内に集約するとよい

```
$ go test -p 1 -v ./pkg/...
=== RUN   TestSample11
TestSample11 2021-12-12 16:25:45.602542054 +0900 JST m=+0.000403602
TestSample11 End 2021-12-12 16:25:46.6030248 +0900 JST m=+1.000886478
--- PASS: TestSample11 (1.00s)
=== RUN   TestSample12
TestSample12 2021-12-12 16:25:46.603161328 +0900 JST m=+1.001022956
TestSample12 End 2021-12-12 16:25:47.60327697 +0900 JST m=+2.001138588
--- PASS: TestSample12 (1.00s)
=== RUN   TestSample13
TestSample13 2021-12-12 16:25:47.603402918 +0900 JST m=+2.001264536
TestSample13 End 2021-12-12 16:25:48.603847582 +0900 JST m=+3.001709301
--- PASS: TestSample13 (1.00s)
PASS
ok      github.com/syunkitada/go-samples/test-sample/parallel/pkg/sample1       3.004s
=== RUN   TestSample11
TestSample11 2021-12-12 16:25:48.741053013 +0900 JST m=+0.000543425
TestSample11 End 2021-12-12 16:25:49.741693247 +0900 JST m=+1.001183760
--- PASS: TestSample11 (1.00s)
=== RUN   TestSample12
TestSample12 2021-12-12 16:25:49.741836468 +0900 JST m=+1.001326880
TestSample12 End 2021-12-12 16:25:50.741949051 +0900 JST m=+2.001439504
--- PASS: TestSample12 (1.00s)
=== RUN   TestSample13
TestSample13 2021-12-12 16:25:50.74205425 +0900 JST m=+2.001544672
TestSample13 End 2021-12-12 16:25:51.74250719 +0900 JST m=+3.001997703
--- PASS: TestSample13 (1.00s)
PASS
ok      github.com/syunkitada/go-samples/test-sample/parallel/pkg/sample2       3.004s
```

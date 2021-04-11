# Testing

## Godoc

- Look godoc on Godoc with github
- Access to https://godoc.org/github.com/syunkitada/go-samples/pkg/test_example/simple

- Serve godoc on localhost
- Access to http://192.168.10.103:6060/pkg/github.com/syunkitada/go-samples/pkg/test_example/simple/

```
# Serve godoc
$ godoc -http=0.0.0.0:6060
```

- Look godoc on cli

```
$ go doc github.com/syunkitada/go-samples/pkg/test_example/simple
```

## Basic Test

```
# Test
$ go test simple/simple_test.go -v
=== RUN   TestHello
--- PASS: TestHello (0.00s)
PASS
ok      command-line-arguments  0.001s


# Test package
$ go test github.com/syunkitada/go-samples/test-sample/simple
ok      github.com/syunkitada/go-samples/test-sample/simple     0.001s
```

## Test Example Code

```
$ go test simple/simple_example_test.go -v
=== RUN   Example
--- PASS: Example (0.00s)
=== RUN   ExampleHello
--- PASS: ExampleHello (0.00s)
=== RUN   ExampleFoo
--- PASS: ExampleFoo (0.00s)
=== RUN   ExampleFoo_Hello
--- PASS: ExampleFoo_Hello (0.00s)
=== RUN   ExampleFoo_Hello_world
--- PASS: ExampleFoo_Hello_world (0.00s)
PASS
ok      command-line-arguments  (cached)
```

## Test Coverage

通常の Coverage はテストを書いたパッケージのみで計算されるため、パッケージ全体の Coverage が出せない

```
$ go test -coverprofile=coverage.out ./coverapp/...
?       github.com/syunkitada/go-samples/test-sample/simple2    [no test files]
ok      github.com/syunkitada/go-samples/test-sample/simple2/api        0.009s  coverage: 100.0% of statements
?       github.com/syunkitada/go-samples/test-sample/simple2/lib        [no test files]

$ go tool cover -func=coverage.out
github.com/syunkitada/go-samples/test-sample/simple2/api/api.go:7:      GetResult       100.0%
total:                                                                  (statements)    100.0%
```

以下のようにすると、全体の Coverage が出せる

また、lib パッケージのテストは書いてないがテストコードから間接的に呼び出しているので 100%となる

```
$ go test --coverpkg ./coverapp/... -coverprofile=coverage.out ./coverapp/...
?       github.com/syunkitada/go-samples/test-sample/coverapp   [no test files]
ok      github.com/syunkitada/go-samples/test-sample/coverapp/api       0.007s  coverage: 66.7% of statements in ./coverapp/...
?       github.com/syunkitada/go-samples/test-sample/coverapp/lib       [no test files]

$ go tool cover -func=coverage.out
github.com/syunkitada/go-samples/test-sample/coverapp/api/api.go:5:     GetResult       100.0%
github.com/syunkitada/go-samples/test-sample/coverapp/lib/lib.go:3:     Add             100.0%
github.com/syunkitada/go-samples/test-sample/coverapp/main.go:9:        main            0.0%
total:                                                                  (statements)    66.7%
```

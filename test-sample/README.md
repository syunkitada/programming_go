# Testing


## Godoc
* Look godoc on Godoc with github
* Access to https://godoc.org/github.com/syunkitada/go-samples/pkg/test_example/simple

* Serve godoc on localhost
* Access to http://192.168.10.103:6060/pkg/github.com/syunkitada/go-samples/pkg/test_example/simple/

```
# Serve godoc
$ godoc -http=0.0.0.0:6060
```

* Look godoc on cli

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

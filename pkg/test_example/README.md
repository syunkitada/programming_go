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
$ go test pkg/test_example/simple/simple_test.go
ok      command-line-arguments  (cached)

# Test package
$ go test github.com/syunkitada/go-samples/pkg/test_example/simple
```


## Test Example Code
```
$ go test pkg/test_example/simple/simple_example_test.go -v
```

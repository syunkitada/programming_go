# godo

- https://github.com/go-godo/godo
- go 用のタスクランナーです
- ファイルの変更を検知して、shell を実行したり、go プログラムを再起動できます

## 例 1

```
package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	p.Task("test", nil, func(c *do.Context) {
		c.Bash("go test ./pkg/...")
	}).Src("pkg/**/*.go")

	p.Task("loop-sample", nil, func(c *do.Context) {
		c.Start("main.go", do.M{"$in": "cmd/loop-sample"})
	}).Src("cmd/loop_sample/*.go", "pkg/loop_sample/**/*.go")
}

func main() {
	do.Godo(tasks)
}
```

```
# ファイル変更を検知してgo testを実行します
$ go run cmd/godo/main.go test --watch

# ファイル変更を検知してloop-sampleを再実行します
$ go run cmd/godo-sample/main.go loop-sample --watch
```

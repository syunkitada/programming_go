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

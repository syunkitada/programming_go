package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	p.Task("sample-api", nil, func(c *do.Context) {
		c.Start("main.go --port 8080", do.M{"$in": "cmd/sample-api"})
	}).Src("cmd/sample-api/*.go", "pkg/**/*.go")
}

func main() {
	do.Godo(tasks)
}

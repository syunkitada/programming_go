package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	p.Task("helloworld", nil, func(c *do.Context) {
		c.Start("main.go", do.M{"$in": "helloworld"})
	}).Src("helloworld/**/*.go")
}

func main() {
	do.Godo(tasks)
}

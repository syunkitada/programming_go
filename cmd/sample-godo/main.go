package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	p.Task("compile-pb", nil, func(c *do.Context) {
		c.Bash("protoc -I pkg/grpc/simple/pb pkg/grpc/simple/pb/pb.proto --go_out=plugins=grpc:pkg/grpc/simple/pb")
	}).Src("pkg/grpc/**/*.proto")

	p.Task("sample-grpc-simple-server", nil, func(c *do.Context) {
		c.Start("main.go", do.M{"$in": "cmd/sample-grpc-simple-server"})
	}).Src("pkg/grpc/**/*.go")

	p.Task("sample-grpc-simple-client", nil, func(c *do.Context) {
		c.Start("main.go", do.M{"$in": "cmd/sample-grpc-simple-client"})
	}).Src("pkg/grpc/**/*.go")
}

func main() {
	do.Godo(tasks)
}

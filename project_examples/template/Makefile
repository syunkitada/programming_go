gen:
	oapi-codegen -generate "types" -package oapi api/@appname/spec.yaml > api/@appname/oapi/types.gen.go
	oapi-codegen -generate "client" -package oapi api/@appname/spec.yaml > api/@appname/oapi/client.gen.go
	oapi-codegen -generate "server" -package oapi api/@appname/spec.yaml > api/@appname/oapi/server.gen.go
	oapi-codegen -generate "spec" -package oapi api/@appname/spec.yaml > api/@appname/oapi/spec.gen.go
	go generate ./internal/...

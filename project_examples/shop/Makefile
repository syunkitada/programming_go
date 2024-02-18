gen:
	oapi-codegen -generate "types" -package oapi api/shop-api/spec.yaml > api/shop-api/oapi/types.gen.go
	oapi-codegen -generate "client" -package oapi api/shop-api/spec.yaml > api/shop-api/oapi/client.gen.go
	oapi-codegen -generate "server" -package oapi api/shop-api/spec.yaml > api/shop-api/oapi/server.gen.go
	oapi-codegen -generate "spec" -package oapi api/shop-api/spec.yaml > api/shop-api/oapi/spec.gen.go

	go generate ./internal/...

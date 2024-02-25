# oapi-codegen

```
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.3.8
```

```
$ vim spec.yaml
```

- 型定義
- http client
- http server
- OpenAPI spec

```
oapi-codegen -generate "types" -package openapi spec.yaml > ./openapi/types.gen.go
oapi-codegen -generate "client" -package openapi spec.yaml > ./openapi/client.gen.go
oapi-codegen -generate "server" -package openapi spec.yaml > ./openapi/server.gen.go
oapi-codegen -generate "spec" -package openapi spec.yaml > ./openapi/spec.gen.go
```

以下のコマンドで 1 ファイルで生成する方法もある
これは好みで選ぶと良いが、個人的にはファイルが分割されてるほうが好み

```
oapi-codegen -package openapi -package openapi spec.yaml > ./openapi/openapi.gen.go
```

https://github.com/deepmap/oapi-codegen/blob/master/examples/petstore-expanded/internal/config.yaml

directory structure

```
cmd/
  app-name/
    main.go
pkg/
  db-api
  db-model
  app-name/
    config/
    api/
      handler/
      middleware/
      openapi/  # generated codes
      spec.yaml
  lib/
    config
    logger
scripts/
  make-env.sh
  make-pkg.sh
tests/
  e2e/
```

# OpenAPI

- OpenApi
  - https://github.com/OAI/OpenAPI-Specification
  - プログラミング言語に依存しない API の仕様書を定型化したもの
  - OpenAPI に準拠した仕様書を YAML, JSON で書いて、それを元にジェネレータでコードを自動生成して利用する
  - コードから YAML、JSON に変換するパターンもある
    - コードの規模が増えると単一 YAML からコード生成するよりも、コードから逆生成するパターンのほうが多い（気がする）
- Swagger
  - https://swagger.io/
  - OpenAPI のためのツール類
    - Swagger Editor
      - ブラウザ上で API 仕様書を書くためのエディタ
    - Swagger UI
      - API 仕様書からドキュメントを生成するツール
    - Swagger Codegen
      - API 仕様書からコードを生成するツール
  - API 仕様書が、2.0 では Swagger Specification だったが、3.0 では OpenAPI Specification になった？
  - go-swagger
    - https://github.com/go-swagger/go-swagger
    - Swagger2.0 に対応(3.0 は非対応)
  - API ドキュメントとして Swagger UI だけ使うケースが多い気がする
- ライブラリ
  - go-restful
    - https://github.com/emicklei/go-restful
    - k8s で使われてるやつ
  - go-swagger
    - https://github.com/go-swagger/go-swagger
  - open-go
    - https://github.com/ogen-go/ogen
  - oapi-codegen
    - https://github.com/deepmap/oapi-codegen
  - kin-openapi
    - https://github.com/getkin/kin-openapi

## swagger-ui

```
$ sudo docker run --name swagger --rm -p 3000:8080 -e SWAGGER_JSON=/swagger/swagger.json -v $PWD/lib_go-restful/sample1:/swagger swaggerapi/swagger-ui

$ sudo docker run --name swagger --rm -p 3000:8080 -e SWAGGER_JSON=/swagger/swagger.json -v $PWD/lib_go-swagger/sample1:/swagger swaggerapi/swagger-ui
```

## メモ

- [Go における API ドキュメントベースの Web API 開発について登壇しました](https://future-architect.github.io/articles/20210427c/)
  - ソースコードから vs API ドキュメントから
    - API ドキュメントからのがよい
  - どれを使うのが良いか？
    - OpenAPI 2.0 でよいなら go-swagger
    - OpenAPI 3.0 がよいなら opai-codegen + Echo

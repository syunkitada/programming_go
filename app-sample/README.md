# app-sample

- swagger でのアプリケーション作成のサンプルです

## swagger によりアプリケーションのひな型を作ります

- generate server
  - --main-package
    - エントリーポイントの main.go を生成する場所を指定します
    - --exclude-main
      - main.go を生成したくない場合は、これを指定します
  - cmd

```
mkdir -p pkg/sample-api/
vim pkg/sample-api/swagger.yml

swagger generate server \
  --name sample-api \
  --main-package sample-api \
  --spec pkg/sample-api/swagger.yml \
  --server-package pkg/sample-api/gen/api \
  --model-package pkg/sample-api/gen/models \
  --strict-additional-properties

swagger generate client \
  --name sample-api \
  --spec pkg/sample-api/swagger.yml \
  --model-package pkg/sample-api/gen/models \
  --client-package pkg/sample-api/gen/client \
  --strict-additional-properties


# cliをgenerateする
swagger generate cli \
  --name sample-api \
  --spec pkg/sample-api/swagger.yml \
  --target pkg/sample-api/gen \
  --strict-additional-properties

# generate cliでは、mainパッケージの場所を指定できないので手でmain パッケージをcmd配下にコピーする
cp -r pkg/sample-api/cmd/cli cmd/sample-api

go get -u github.com/spf13/cobra
```

コードの自動生成後は gen/api/configure_xxx.go のみを編集してアプリケーションを開発していきます
それ以外の自動生成されたコードは.gitignore に入れます

```
# ignore generated files by swagger-go
cmd/*
!cmd/godo
!cmd/sample-cli
pkg/sample-api/gen/*
!pkg/sample-api/gen/api
pkg/sample-api/gen/api/*
!pkg/sample-api/gen/api/configure_sample_api.go
```

## godo によりアプリケーションを自動リロードできるようにする

以下のような、 cmd/godo/main.go を作成します。

```
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
```

```
$ go run cmd/godo/main.go sample-api --watch
sample-api rebuilding with -a to ensure clean build (might take awhile)
sample-api 9306ms
sample-api watching /home/owner/programming_go/app-sample/cmd/sample-api
sample-api watching /home/owner/programming_go/app-sample/pkg
2022/11/23 16:46:00 Serving sample API at http://127.0.0.1:8080
```

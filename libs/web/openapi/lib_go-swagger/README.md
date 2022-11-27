# go-swagger

- swagger-codegen はフラットも出るしかサポートしておらず、生成するのはクライアントコードだけで、サーバ用のコードもほとんどがスタブです
- go-swagger は、swagger specification からサーバコード、クライアントコード、CLI コードを生成できる
- ソースから spec を逆生成することもできる

```
# cliツールのインストール
$ go install github.com/go-swagger/go-swagger/cmd/swagger@latest

$ swagger -h
```

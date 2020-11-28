# mock1

- mockgen を利用する

```
$ go get github.com/golang/mock/mockgen
```

## 使い方

- 対象コードに Interface を定義しておく
- mockgen -source program.go -destination mock_program.go

```
# 実行例
$ mockgen -source client/client.go -destination client/mock_client.go
```

## go generate から mockgen を使う

- 一般的には、mockgen を直接実行するのではなく、go generate を使う
- go generate とは
  - コマンドラインから go generate ./... と実行すると
  - 対象コードから //go:generate から始まるコメント行を検索し //go:generate command argument... というコメント行をそのまま実行してくれる
  - 任意のコマンドを実行できるが Go のコードを生成することを想定して作られている
- 対象コードの一番上に以下のコメント行を書いておき、go generate を実行する
  - //go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=\$GOPACKAGE

```
# 実行例
$ go generate ./client/...
```

## test を実行してみる

```
$ go test ./...
ok      github.com/syunkitada/go-samples/test-sample/mock1      0.010s
?       github.com/syunkitada/go-samples/test-sample/mock1/client       [no test files]
```

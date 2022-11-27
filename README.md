# Go

- 勉強用の Go プログラム群です

## コンテンツ

| Link                                        | Description                                            |
| ------------------------------------------- | ------------------------------------------------------ |
| [Getting Started](basic/getting_started.md) | 入門                                                   |
| [基礎](basic/README.md)                     | 基礎                                                   |
| [ライブラリ勉強用](libs/README.md)          | ライブラリ勉強用のメモ書きやサンプルコードなどの置き場 |
| [チップス](tips/README.md)                  | 雑多なプログラム群                                     |

## 最低限のルール

#### プロジェクト作成時

- go mod init は、以下のようにサブプロジェクトごとに初期化して利用します
- 基本的に外部からの利用を想定しないのでディレクトリ名で初期化します

```
mkdir libs/lib_project1
cd libs/lib_project1
go mod init lib_project1
```

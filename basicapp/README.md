# basicapp
* アプリケーションの基本テンプレート


## ライブラリ メモ
* CLI管理: https://github.com/spf13/cobra
    * 既存のファイルをいじらずにモジュール的に機能を追加できる
    * サブコマンドごとのファイル分割、ディレクトリ分割が用意にできる
    * github.com/urfave/cliというのも有名だが、サブコマンドの管理がやりずらいのでcobraのほうが良さげ
* 設定ファイル管理: https://github.com/BurntSushi/toml
    * toml設定ファイル管理用のライブラリ
    * 設定値をstructで扱え、デフォルト値も(コマンド引数などによって)動的に設定できる
* ロギング: https://github.com/google/glog
    * Googleのloggingライブラリ


# Go Run
```
$ go run cmd/grcp-sample/main.go -h
```


# テスト実行
```
$ go test pkg/ctl/main_test.go
ok      command-line-arguments  0.001s
```


# パッケージ作成
```
$ make rpm
```

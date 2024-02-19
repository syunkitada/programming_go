# Go のプロジェクトレイアウト

go のプロジェクトレイアウトは多岐にわたっており、悩ませる要因の一つである。

一般的に出回ってるレイアウトには以下がある。

- [project-layout](https://github.com/golang-standards/project-layout)
  - 標準的なレイアウトの一例として有名な奴
  - これを日本語で解説してくれてるやつ
    - https://qiita.com/vengavengavnega/items/2235589445dd0effda05

プロジェクトの種別によって構成も変わるので一概にこれがよいというのは難しい。

ここでは、k8s のようなマルチアプリケーションを管理するプロジェクトを想定し、個人的にこうするとよさそうという構成を乗せる。

```
api/
  [application name]/    # OpenAPIの定義書
    openapi.yaml
cmd/
  [application name]/    # 各アプリケーションのエントリポイント
    main.go
internal/                # 各アプリケーションの非公開実装
  [application name]/
    handler/             # OpenAPIのハンドラ実装(application層、usecase層などと呼ばれる、ビジネスロジックをここに置いてはいけない)
    controller/          # ビジネスロジックの実装
  lib/ or util/          # 共通ライブラリ
    ...
  model/                 # モデル定義
    ...
  repository             # 永続化データを操作するためのモジュール
    ...
pkg/                     # 公開したい実装（オプショナル）
  ...
scripts/ or tools/       # 環境構築などに使われる雑多なスクリプト郡やツール群(tools派よりscripts派が多い印象）
  ...
build/ or ci/            # ビルド系のツールなど（オプショナル）（ci派よりbuild派が多い印象）
tests/                   # unitテスト以外のテストをここに置く（オプショナル）
  e2e/
examples/                # 公開したライブラリの実装の利用例をここに置く（オプショナル）
  ...
```

- internal について
  - 昔は internal, pkg などの区別はなく、pkg に本来 internal であるべきものもごっちゃにしていた
    - pkg というのも誰が始めたかわからないが、有名な k8s などがそうなので、みなそれに習ってるだけのように見える
  - go1.4 では[Internal packages](https://go.dev/doc/go1.4#internalpackages)というのが導入され、internal と名のついたディレクトリは非公開となる
  - このため公開する必要のないものは internal に置くのが主流となっている
- repository について
  - Repository パターンを採用している
  - Repository は、永続化データを操作するためのモジュールで、DB 接続や SQL クエリなどを隠ぺいします
  - 方針
    - 機能や役割に応じて File は分けても Repository は分けない
      - Repository を分けてしまうと、あるデータのアクセスのためにどの Repository を使っていいかわからなくなる
    - File の分割は塊を意識する
      - Model の 1 ファイルと、Repository の 1 ファイルが一致してるとわかりやすい
    - Repository にビジネスロジックを入れない
      - データ処理中にビジネスロジックが必要な場合は、ライブラリ呼び出しなどして外だし、ビジネスロジックだけでテストできるようにする
    - 不用意にトランザクションに頼らない（しょうがない時はある）

# grpc

- https://grpc.io/docs/languages/go/quickstart/

## 類似ライブラリ

- gogo
  - golang/protobuf のフォークしたもの
  - 本家よりもパフォーマンスが良い？
  - メンテ状況もちょっとあやしい？
  - 基本的には本家でよいと思う
- ttrpc
  - https://github.com/containerd/ttrpc
  - containerd で利用されている grpc 互換のプロトコル
  - メモリのフットプリントを軽くしたいなら採用を検討するとよいと思う

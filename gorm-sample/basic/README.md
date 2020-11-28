# 基本的な使い方

## 更新処理について

- .Updates(map[string]ingerface{})で部分更新するのがよい
  - .Updates(map[string]interface{}{"name": "hello", "age": 18})
- Struct を利用した更新は予期しない動作をする場合があるため利用は避ける
  - Struct を利用した Save は、これは空値も含めてすべてのカラムを一律で更新する
    - 予期しないカラムを更新する可能性がある
  - Struct を利用した Updates は空値が除外されて、更新される
    - 0 や null で更新したい場合でも無視される

## Auto Migration について

- http://gorm.io/docs/migration.html
- Auto Migration:
- WARNING: AutoMigrate will ONLY create tables, missing columns and missing indexes, and WON’T change existing column’s type or delete unused columns to protect your data.
- カラムの作成とインデックスの追加のみ行い、カラムの変更や、削除は行わない
- 本番環境のカラムの変更をする場合には、手動で変更する必要がある
- タグについて
  - http://gorm.io/ja_JP/docs/models.html
  - type に明示的にカラム定義を書いておくと、CREATE TABLE 時にそのまま使われる
  - 複数 primaly_key を行う場合は、primary_key タグを利用する

## gorm.Model の利用について

- deleted_at に timestamp を利用する論理削除前提のモデルであることに注意する
  - 論理削除が不要な場合は無駄に計算コストが増えるので利用しないほうがよい

## primary key について

- primary key に int unsigned を利用した場合
  - 最大値は 4294967295 (4 byte: 43 億 程度)
  - 毎秒 1 件追加すると、ざっくり 100 年で消費できる
- primary key に bigint(unsigned)を利用した場合
  - 最大値は、18446744073709551615 (8 byte: 1844 京 6744 兆 737 億)
  - 毎秒 584942 件追加すると 1,000,000 年で消費できる
  - 世界人口 100 億人と仮定すると、全員が毎秒投稿しても 58 年かかる
- primary key に 文字列を利用した場合

## join について

- string を使用しての join
- int を使用しての join

## 型について

- text 型について

# 論理削除とユニーク

## deleted フラグを利用する方法

- deleted はプライマリーキーと同じ型にしておく
- deleted は NOT NULL とする
- ユニークにしたいフィールドと、deleted で複合のユニークインデックスを作成する
- 論理削除時に deleted にプライマリーキーと同じ値を保存する（これで削除されたとみなす

## exist フラグを利用する方法(NULL の特性を利用する)

- exist を用意する
  - bool(tiny int), NOT NULL, default true(1)
- ユニークにしたいフィールドと、exist で複合ユニークインデックスを作成する
- 論理削除時に exist に NULL をセットする
- 複合ユニークインデックスに NULL があるとユニークの制約が外れる

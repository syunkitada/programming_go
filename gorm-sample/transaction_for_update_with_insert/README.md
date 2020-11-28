# FOR UPDATE を Insert のために使ってみる

- Insert しようとしてるテーブルに対して Select ... FOR UPDATE を実行してロックする
- そのトランザクションで Insert すると、待機してた別トランザクションは再開時に、以下のエラーを出して終了する
  - Deadlock found when trying to get lock; try restarting transaction
- Insert のために FOR UPDATE を実行するのはよろしくない?

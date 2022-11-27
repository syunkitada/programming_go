# Transaction For Update

- [Locking Reads](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking-reads.html)
- Select .. For Update
  - その行のインデックスエントリに関してロックする
  - 他のトランザクションはそのインデックスエントリを使った Select や Update を実行するとそこで待たされる
- Select .. For Share
  - For Update は Select も待たされるが、For Share は Select は他のトランザクションも実行できる

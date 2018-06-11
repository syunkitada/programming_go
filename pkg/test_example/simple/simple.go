// 先頭部分には、Copyrightを書くのが一般的
// package文の直上のコメントが、godocで表示されるため、この部分はgodocには表示されない

// This is test sample package.
// ここにパッケージの概要を書く。
//
// Hoge
//
// 空行を上下に入れると、h3になるので、
// 細かい題目はここように区切って説明を記述する。
//
// また、改行しても空行を入れない限り、一つの<p>で扱われるため、文を区切る場合はこのように空行を入れる。
//
// Piyo
//
// 以下のように先頭にスペースを入れることで、一つの<pre>で扱われる。
//   hello
//   test
//
package simple

// This return "hello".
//
// 関数のgodocは、関数宣言の直上に記述する。
//
func Hello() string {
	return "hello"
}

// This is foo.
//
// structのgodocは、struct宣言の直上に記述する。
type Foo struct {
	Name string
}

// This return "hello " + Name
//
// メソッドのgodocは、メソッド宣言の直上に記述する。
func (f *Foo) Hello() string {
	return "hello " + f.Name
}

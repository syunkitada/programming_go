package simple_test

import (
	"fmt"
	"github.com/syunkitada/go-sample/pkg/test_example/simple"
)

func Example() {
	fmt.Println(simple.Hello())
	// Output:
	// hello
}

func ExampleHello() {
	fmt.Println(simple.Hello())
	// Output:
	// hello
}

func ExampleFoo() {
	f := simple.Foo{Name: "example"}
	fmt.Println(f.Hello())
	// Output:
	// hello example
}

func ExampleFoo_Hello() {
	f := simple.Foo{Name: "example"}
	fmt.Println(f.Hello())
	// Output:
	// hello example
}

func ExampleFoo_Hello_world() {
	f := simple.Foo{Name: "world"}
	fmt.Println(f.Hello())
	// Output:
	// hello world
}

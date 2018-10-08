package tracing

import (
	"fmt"
)

func Main() {
	fmt.Print("hoge")
	test()
}

func test() {
	span := opentracing.StartSpan("test")
	defer span.Finish()

	fmt.Print("hogeaa")
}

func childTest(parentSpan opentracing.Span) {
	sp := opentracing.StartSpan(
		"test",
		opentracing.ChildOf(parentSpan.Context()))
	defer sp.Finish()
	fmt.Print("hogeaa child")
}

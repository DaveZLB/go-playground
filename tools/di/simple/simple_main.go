package main

import (
	"fmt"
	"github.com/go-playground/tools/di/simple/di"
)

type foo struct {
	FooMessage string
}

func NewFoo(m string) *foo {
	return &foo{
		FooMessage: m,
	}
}

type bar struct {
	BarMessage string
	Foo        *foo
}

func NewBar(m string, foo *foo) *bar {
	return &bar{
		BarMessage: m,
		Foo:        foo,
	}
}

func main() {
	container := di.NewContainer(
		di.ServiceConstructorMap{
			"foo": func(get di.Get) interface{} {
				return NewFoo("fooMessage")
			},
			"bar": func(get di.Get) interface{} {
				return NewBar("barMessage", get("foo").(*foo))
			},
		})

	b := container.Get("bar").(*bar)
	fmt.Println(b.BarMessage)
	fmt.Println(b.Foo.FooMessage)
}

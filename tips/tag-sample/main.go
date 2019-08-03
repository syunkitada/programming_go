package main

import (
	"fmt"
	"reflect"
)

type Hoge struct {
	Name string `label1:"name:hoge;size:10;" label2:"name:piyo"`
}

func main() {
	hoge := Hoge{
		Name: "test",
	}
	cmdType := reflect.TypeOf(hoge)
	fmt.Println(cmdType.Name())
	// fmt.Println(reflect.TypeOf(hoge).Field(0).Tag)
	// fmt.Println(reflect.TypeOf(hoge).Field(0).Tag.Get("label1"))
	// fmt.Println(reflect.TypeOf(hoge).Field(0).Tag.Get("label2"))
	// fmt.Println(reflect.TypeOf(hoge).Name())
}

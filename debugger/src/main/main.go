package main

import (
	"fmt"
	"reflect"
	"rules"
	"runtime"
)

func main() {
	fmt.Println(reflect.TypeOf(rules.GetA()))
	runtime.Breakpoint()
	fmt.Println(reflect.ValueOf(rules.GetA))
}

package main

import (
	"fmt"
	"sub"
	"os"
)

var _ = sub.Glob

func main() {
	if len(os.Args)>0 {
		fmt.Println("Hi")
	}
}
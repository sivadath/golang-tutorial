package main

import (
	"fmt"
	"functions"
	"flag"
	"os"
)

func main() {
	var booleanFlg bool
	flag.BoolVar(&booleanFlg,"bool",true, "bool flag enables boolean functionality.")
	strFlg := flag.String("string","nil","string falg enables string functionality. set it to apple to see apple")

	flag.Parse()


	if *strFlg != "nil" {
		fmt.Println("string falg set to:", *strFlg)
	}
	fmt.Println("Boolean falg set to:", booleanFlg)

	if err, hostName := os.Hostname(); err == "" {
		fmt.Println("Host name of current device obtained from os package: ", hostName)

	}

	fmt.Println("Arguments obtained using os package:", os.Args)

	fmt.Println("Pid of obtained from os package:", os.Getpid())

	fmt.Println("Outcome of an inline function Add with arguments 1, 2: ", functions.Add(1,2))
	fmt.Println("Outcome of a variadic function AddMultiple with arguments 1, 2, 3, 4: ", functions.AddMultiple(1,2, 3, 4))
	functions.AnonymousFunction()
	functions.Articles("himalayas")
	functions.CallBack(functions.AnonymousFunction)
	functions.Defer()
	emp := functions.Employee{"imran", 7}
	functions.ModifyValue(emp)
	fmt.Println("After passing value to ModifyValue, value of emp:", emp)
	ptr := &emp
	functions.Modify(ptr)
	fmt.Println("After passing pointer to Modify, value of emp:", emp)
	//while implementation:
	var i int
	for ; i != 10; i++ {

	}
}

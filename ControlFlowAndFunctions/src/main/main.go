package main

import (
	"flag"
	"fmt"
	"functions"
	"network/HTTP"
	"strconv"
)

func main() {
	flag.Parse()
	switch functions.Fun {
	case "inline" :
		var a,b int
		fmt.Println("Give two numbers:")
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Println("sum of two numbers given:", functions.Add(a,b))
		fmt.Println(`To see what functions were inlined, use build command "go build -gcflags -m main.go"`)
	case "anonymous" :
		functions.AnonymousFunction()
	case "variadic" :
		fmt.Println("Give multiple number to be added and enter q to stop:")
		var arr []int
		for {//Infinite for loop implementation.
			var a string
			fmt.Scan(&a)
			if a == "q" {
				break
			}
			if i, err := strconv.Atoi(a); err == nil {
				arr = append(arr, i)
			}
		}
		fmt.Println("Sum of all the numbers given:", functions.AddMultiple(0, arr...))
	case "callback" :
		functions.CallBack(functions.AnonymousFunction)
	case "defer" :
		functions.Defer()
	case "expression":
		var noun string
		fmt.Println("Give a noun:")
		fmt.Scan(&noun)
		functions.Articles(noun)
		fmt.Println("Check functions.Articles definition to understand function as expression concept.")
	case "passValue" :
		var emp functions.Employee
		fmt.Println("Give name of the employee:")
		fmt.Scan(&emp.Name)
		fmt.Println("Before passing employee data by passing value:", emp)
		functions.ModifyValue(emp)
		fmt.Println("After passing employee data by passing value (we see number not getting updated):", emp)
	case "passPointer" :
		var emp functions.Employee
		fmt.Println("Give name of the employee:")
		fmt.Scan(&emp.Name)
		fmt.Println("Before passing employee data by passing pointer:", emp)
		functions.Modify(&emp)
		fmt.Println("After passing employee data by passing pointer (we see number got updated):", emp)
	}
	if HTTP.Network == "client" {
		for {
			var ip, port, content string
			fmt.Println("Give ipaddress of http server")
			fmt.Scan(&ip)
			fmt.Println("Give port number of http server")
			fmt.Scan(&port)
			fmt.Println("Give content to be passed to http server")
			fmt.Scan(&content)
			HTTP.HitCaseConverterServer(content, ip, port)
			fmt.Println("Press q to stop any other key to rehit the server")
			fmt.Scan(&content)
			if content == "q" {
				break
			}
		}
	}else if HTTP.Network == "server" {
		var port string
		fmt.Println("Give port number of http server")
		fmt.Scan(&port)
		HTTP.StartCaseConverterServer(port)
	} else if HTTP.Network != "" {
		fmt.Println("Invalid network option given check help to see valid options")
	}
}

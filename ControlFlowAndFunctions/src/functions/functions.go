package functions

import (
	"fmt"
	"reflect"
	"runtime"
)


func init() {
	fmt.Println("init of functions package invoked.")
}

//INLINE FUNCTION

//Add will be made inline function by go compiler can be checked with go build flags go build -gcflags -m main.go
func Add(a, b int)(c int) {
	return a+b
}


//VARIADIC FUNCTION
//AddMultiple is a variadic function it take any number of arguments.
func AddMultiple(base int, adders ...int) int {
	for _, num := range adders {
		base += num
	}
	return base
}


//FUNCTION AS EXPRESSION
//Articles is a sample function that adds articles to given nouns.
func Articles(noun string) {

	//AppendArticle is declared as an expression
	AppendArticle := func(article, noun string) string {
		return article + noun
	}
	if (noun[0] == 'a') || (noun[0] == 'e') || (noun[0] == 'i') || (noun[0] == 'o') || (noun[0] == 'u') {
		fmt.Println(AppendArticle("an ", noun))
	} else if noun == "himalayas"{
		fmt.Println(AppendArticle("the ", noun))
	} else {
		fmt.Println(AppendArticle("a ", noun))
	}
}

//ANONYMOUS FUNCTION
//AnonymousFunction shows the implementation of anonymous function.
func AnonymousFunction() {
	k := func(a string, b string)(i int){
		fmt.Println("Anonymous function called")
		return i
	}
	fmt.Println("Name of the function invoked:",GetFunctionName(k))
}


//CALL BACK FUNCTION
//CallBack function shows implementation of Call back function. Also shows that a function can be returned by another function.
func CallBack(fn func()) (func()) {
	fmt.Println("Received ",GetFunctionName(fn), "as call back")
	fn()
	return fn
}

//DEFER

//Defer shows implementation of defer
func Defer() {
	fmt.Println("Sample outcome of defer function:")
	printStatement := func(count string) {fmt.Println("Defer number ",count)}
	defer printStatement("First inserted defer statement")
	defer printStatement("Second inserted defer statement")
	func () {
		defer printStatement("Defer in nested function one.")
	}()
}


type Employee struct {
	Name string
	Number int
}
//Pass by value using pointer
func Modify(sample *Employee) {
	sample.Number = len(sample.Name)
}

//Pass by value
func ModifyValue(sample Employee) {
	sample.Number = len(sample.Name)
}


func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
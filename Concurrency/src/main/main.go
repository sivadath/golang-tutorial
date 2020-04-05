package main

import (
	"concurrency"
	"errors"
	"fmt"
	_ "runtime"
	_ "sync"
	_ "sync"
	_ "time"
)

type inter interface {
	foo()
	zoo()
}


type strt struct {
	I int
	N string
}

func foo()  (err error) {
	return errors.New("This is a dummy error")
}

func fun1() {
	if err := foo(); err!=nil {
		panic(err.Error())
	}
}

func fun2() {
	defer func() {
		if err := recover();err != nil {
			fmt.Println("We received a panic in fun2")
		}
	}()
	fun1()
	fmt.Println("fun2 executed successfully")
}

func main () {
	concurrency.WaitGroupDemo()
	//ch := make(chan int,0)
	//ch <- 10
	fun2()
}


func receive (ch <-chan int) {
	i := <-ch
	fmt.Println("Received value from channel:",i)
}

func send (ch chan<- int) {
	ch <- 10
}







/*mp := map[string]int{
	"one" : 1,
	"two" : 2,
}
arr := make([]int,10)
for index, value := range arr  {
	fmt.Println(index,value)
}
for key, value := range mp {
	fmt.Println(key, value)
}

if val, exist  := mp["one"]; exist {
	fmt.Println("One is there in map", val)
}else {
	fmt.Println("One is not available in map")
}

if val,exist := mp["three"]; exist {
	fmt.Println("Three is there in map",val)
}else {
	fmt.Println("Three is not available in map")
}
a, _ := foo()
fmt.Println(a)
var wg sync.WaitGroup
wg.Add(2)
ch := make(chan int)

go func() {
	for {//i:= 0; i <10 ; i++ {
		r1, flg := <-ch
		if !flg {
			break
		}
		fmt.Println("Received r1:",r1)
		runtime.Gosched()
	}
	wg.Done()
}()

go func() {
	for i:= 20; i < 30; i++ {
		ch <- i
		runtime.Gosched()
	}
	close(ch)
	wg.Done()
}()
wg.Wait()
fmt.Println("Completed execution")*/


/*ch := make(chan int,2)
ch <- 10
ch <- 11
ch <- 13
fmt.Println("Completed execution",<-ch)*/
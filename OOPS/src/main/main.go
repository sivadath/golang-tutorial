package main

import (
	"fmt"
)

type baser interface {
	PrintBase()
}

type deriveBaser interface {
	baser
	PrintDerive()
}

type base struct {
	I int
	S string
}

func (b base) PrintBase() {
	fmt.Println("Base variable, I, S:",b.I,b.S)
}

type derive struct {
	base
	B bool
}

func (d derive) PrintBase() {
	fmt.Println("I'm print base of derived sruct:.,I,s,b:",d.I,d.S,d.B)
}

type pseudoer interface {
	pseudo()
}

type pseudoImplementer struct {
	pseudoer
}

func PseudoCaller(p pseudoer) {
	p.pseudo()
}


var fn  = func ()  {
	fmt.Println("I'm called")
}

func SingleTon() {
	if fn != nil {
		fn()
		fn = nil
	}
}

type singl struct {
	I int
}

var a = singl{0}

func GetSingleObj() singl {
	return a
}

func main() {
	d := derive{base{1,"one"},true}

	d.base.PrintBase()

	d.PrintBase()
	a := singl{}

//	p := pseudoImplementer{}
	//p.pseudo()

	SingleTon()
	SingleTon()
	SingleTon()
	GetSingleObj()

}

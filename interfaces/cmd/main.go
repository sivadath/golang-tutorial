package cmd

import "fmt"

type polygon struct {
	sides int
	area float64
	perimeter float64
}

type poly interface {
	Area()float64
	Perimeter()float64
}

type Area interface {
	Area()float64
}


type triangle struct {
	base int
	height int
}


func (t triangle)Area()float64 {
	return  float64(t.base * t.base) * 0.5
}


func (t triangle)Perimeter()float64 {
	return float64(2* t.base + t.height)
}

func (t triangle)SidesCount()int {
	return 3
}

type square struct {
	side int
}

func (s square) Area() float64 {
	return float64(s.side*s.side)
}

func (s square) Perimeter() float64 {
	return float64(4 * s.side)
}

func (s square) foo() {
	fmt.Println("foo called")
}



func PrintPoly(poly poly) {
	fmt.Println("area:",poly.Area())
	fmt.Println("perimeter:",poly.Perimeter())
	if tri, ok := poly.(triangle);ok {
		fmt.Println(tri.base)
	}

	if sqr, ok := poly.(square);ok {
		fmt.Println(sqr.side)
	}
}


func Print(i interface{}) {
	fmt.Println(i)
}

func PrintArea(a Area) {
	fmt.Print(a.Area())
}

type rhombus struct{
	area float64
}

func (r rhombus) Area() float64 {
	return r.area
}

func main() {
	t := triangle{height:1,base:2}
	r := rhombus{30.00}
	PrintPoly(t)
	PrintPoly(square{side:2})
	PrintArea(t)
	PrintArea(r)
	Print(t)
}


package generator

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

type arr []int

var a arr


func EvenIterator(a arr) (ch chan int) {
	go func() {
		defer close(ch)
		for _, i := range a{
			if i %2 == 0 {
				ch <- i
			}
		}

	}()
	return ch
}


func main () {
	a = arr{1,2,3,4,5,6,7,8,9}
	for i := range EvenIterator(a) {
		fmt.Println(i)
	}
}

package main

import (
	"fmt"
	"time"
	"sync"
)

type str struct {
	acc sync.RWMutex
	Arr []int
}

var d = str{
	acc: sync.RWMutex{},
	Arr: nil,
}

func (s *str)Add(i int) {
	s.acc.Lock()
	defer s.acc.Unlock()
	s.Arr = append(s.Arr,i)
}

func (s *str)Len() int {
	s.acc.RLock()
	defer s.acc.RUnlock()
	return len(s.Arr)
}

type str2 struct {
	I int
	S string
}

var g = str2 {
	10,
	"dath",
}
var k = false
var m = map[string]string{"one":"okati"}
func main() {
	for i := 0 ;i <8000; i ++ {
		go func() {
			time.Sleep(time.Millisecond*4)
			d.Add(1)
			fmt.Println("After addition len:",d.Len())
			fmt.Println(g)
			fmt.Println(k)
			fmt.Println(m["one"])
		}()
	}
	time.Sleep(time.Second * 10)
}

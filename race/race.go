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

func main() {
	for i := 0 ;i <10; i ++ {
		go func() {
			d.Add(1)
			fmt.Println("After addition len:",d.Len())
		}()
	}
	time.Sleep(time.Second * 10)
}

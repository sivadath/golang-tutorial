package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

func WaitGroupDemo() {
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
	fmt.Println("Completed execution")
}

func ChannelsOnWait(){
	ch := make(chan int,2)
	ch <- 10
	ch <- 11
	ch <- 13
	fmt.Println("Completed execution",<-ch)
}

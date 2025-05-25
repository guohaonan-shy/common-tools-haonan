package _go

import (
	"fmt"
	"sync"
)

func exec() {

	ch1, ch2 := make(chan int), make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			num, ok := <-ch1
			if !ok {
				close(ch2)
				return
			}
			if num > 30 {
				close(ch2)
				return
			}

			fmt.Printf("goroutine 1: %v\n", num)
			ch2 <- num + 1
		}
	}()

	go func() {
		defer wg.Done()
		for {
			num, ok := <-ch2
			if !ok { // channel 已关闭
				close(ch1)
				return
			}
			if num > 30 {
				close(ch1)
				return
			}

			fmt.Printf("goroutine 2: %v\n", num)
			ch1 <- num + 1
		}
	}()

	ch1 <- 1
	wg.Wait()
	return
}

package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int)

func producer() {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("producer: ", i)
	}
	close(ch)
}

func consumer(wg *sync.WaitGroup) {
	// for {
	// 	select {
	// 	case num, ok := <-ch:
	// 		if !ok {
	// 			wg.Done()
	// 			return
	// 		}
	// 		fmt.Println(num)
	// 	}
	// }
	for num := range ch {
		fmt.Println("consumer: ", num)
	}
	wg.Done()
}

func main() {
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go producer()
	go consumer(wg)
	wg.Wait()
}

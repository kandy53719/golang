package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var chA = make(chan struct{})
var chB = make(chan struct{})
var ch = make(chan struct{}, 1)

func PrintA() {
	defer wg.Done()
	for {
		fmt.Println("A")
		chB <- struct{}{}
		<-chA
	}

}

func PrintB() {
	defer wg.Done()
	for {
		<-chB
		fmt.Println("B")
		chA <- struct{}{}
	}

}

func main() {
	wg.Add(2)
	go PrintA()
	go PrintB()
	wg.Wait()
}

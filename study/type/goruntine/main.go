package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func PrintWord(word string, wg *sync.WaitGroup) {
	for i := 1; i < 10; i++ {
		fmt.Println(word)
	}
	wg.Done()
}

func main() {
	fmt.Println("主线程开始")
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go PrintWord("协程1", wg)
	wg.Add(1)
	go PrintWord("协程2", wg)
	wg.Wait()
	fmt.Println("主线程结束")

	//监听输入字符
	buf := bufio.NewReader(os.Stdin)
	b, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}
}

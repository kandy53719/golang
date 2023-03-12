package main

import (
	"fmt"
	"time"
)

func prof() {
	panic("ok")
}

func main() {
	var t = time.NewTicker(time.Second)

	for {
		select {
		case <-t.C:
			//fmt.Println(time.Now())
			go func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Println(time.Now(), err)
					}
				}()
				prof()
			}()
		}
	}
}

package main

import "fmt"

func init() {
	fmt.Println("隐性初始化函数init")
}

func load() {
	fmt.Println("显性初始化函数load")
}

func main() {
	load()
	fmt.Println("主函数main")
	var arr = [3]int{0, 1, 2}
	var s1 = append(arr[0:1], arr[2:]...)
	fmt.Println(arr[0:1], arr[2:], s1)
}

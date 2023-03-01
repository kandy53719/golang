package main

import "fmt"

//泛型函数
func PrintList[T any](list []T) {
	for _, v := range list {
		fmt.Println(v)
	}
}

//泛型类型
type list[T any] []T

func main() {
	PrintList([]int{1, 2, 3})
	PrintList([]string{"张三", "李四", "王五"})

	PrintList(list[int]{1, 2, 3})
}

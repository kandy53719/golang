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

//自定义类型
type CusType interface {
	int | float64 | string
}

// 原始求和
func Sum(m map[string]int) (sum int) {
	for _, v := range m {
		sum += v
	}
	return
}

// 泛型求和
func SumAny[K comparable, V CusType](m map[K]V) (sum V) {
	for _, v := range m {
		sum += v
	}
	return
}

func main() {
	PrintList([]int{1, 2, 3})
	PrintList([]string{"张三", "李四", "王五"})

	PrintList(list[int]{1, 2, 3})
}

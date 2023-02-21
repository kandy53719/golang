package main

import "fmt"

func Add(a int, b int) (num int) {
	return a + b
}

func main() {
	fmt.Println("这是测试", Add(1, 2))
	arr := []int{5, 3, 2, 7, 6}
	InsertionSort(arr)
	//fmt.Println(arr)
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
		fmt.Println(key, j, arr)
	}
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	left := []int{}
	right := []int{}

	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func test() {
	fmt.Println("测试")
}

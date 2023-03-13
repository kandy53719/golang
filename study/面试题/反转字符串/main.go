package main

import "fmt"

func reverse(str string) string {
	var s = []rune(str)
	fmt.Println(s, len(s))
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		fmt.Println(i, s[i], j, s[j])
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func main() {
	fmt.Println(reverse("reverse_ASD中国"))
	//s := []rune("reverse_ASD中国")
	//l := []int
	// for i, v := range s {
	// 	fmt.Printf("i = %v, v = %c \n", i, v)
	// }
}

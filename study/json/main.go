package main

import (
	"encoding/json"
	"fmt"
)

type stu struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var stu1 = stu{Name: "张三", Age: 18}
	bytes, _ := json.Marshal(stu1)
	fmt.Printf("bytes: %v\n", bytes)
	fmt.Printf("bytes: %s\n", bytes)

	var stu2 = new(stu)
	json.Unmarshal(bytes, stu2)
	fmt.Printf("stu2: %v\n", stu2)
}

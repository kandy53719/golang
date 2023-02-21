package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("开始测试")
	m.Run()
	fmt.Println("结束测试")
}

func Test_Add(t *testing.T) {
	result := Add(1, 2)
	t.Logf("result: %v", result)
}

func TestAll(t *testing.T) {
	var tests = []struct {
		a int
		b int
		r int
	}{{1, 2, 3}, {4, 5, 6}}
	for _, test := range tests {
		if Add(test.a, test.b) != test.r {
			t.Logf("失败 参数(%v, %v) 结果%v", test.a, test.b, test.r)
		}
	}
}

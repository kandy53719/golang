package main

// 测试文件以_test.go结尾

import (
	"fmt"
	"testing"
)

// 运行在所有测试之前 负责一些准备操作
func TestMain(m *testing.M) {
	fmt.Println("开始测试")
	m.Run()
	fmt.Println("结束测试")
}

// 单元测试 以Test开发 函数名首字母大写 使用testing.T
func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		// Fatal 致命错误并结束测试
		t.Fatalf("Add 参数：1 2，预计：3，结果：%v ", result)
	} else {
		// Log 输出日志并结束测试
		t.Logf("result: %v", result)
	}
}

// 基准测试 以Benchmark开头
func BenchmarkAdd(b *testing.B) {
	// 基准测试一般用循环来测试 测试次数b.N由程序控制 默认基准测试完成时间为1s
	// BenchmarkAdd-4   	1000000000	         0.5296 ns/op	       0 B/op	       0 allocs/op
	// 1000000000 为循环次数 0.5296 ns/op为单次执行时间 0 B/op 为单次分配内存 0 allocs/op 内存分配次数
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

// 模糊测试
// go test -fuzz=Fuzz -fuzztime 10s
func FuzzAdd(f *testing.F) {
	var testCases = []struct {
		arg1 int
		arg2 int
		want int
	}{
		{1, 2, 3}, {3, 4, 7}, {5, 6, 11},
	}
	// 添加测试集 系统会根据提供的测试用例自动生成测试用例
	for _, val := range testCases {
		f.Add(val.arg1, val.arg2, val.want)
	}
	f.Fuzz(func(t *testing.T, a, b, c int) {
		var result = Add(a, b)
		var want = a + b
		if result != want {
			t.Errorf("Add 结果不符合预期，%c", want)
		}
	})
}

// func TestAll(t *testing.T) {
// 	var tests = []struct {
// 		a int
// 		b int
// 		r int
// 	}{{1, 2, 3}, {4, 5, 6}}
// 	for _, test := range tests {
// 		if Add(test.a, test.b) != test.r {
// 			t.Logf("失败 参数(%v, %v) 结果%v", test.a, test.b, test.r)
// 		}
// 	}
// }

package main

import "fmt"

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Adds(3, 4))
	fmt.Println(Swap(4, 7))
	fmt.Println(Sum(1, 2, 3, 4, 5, 6))
}

// 具名函数
func Add(a, b int) int {
	return a + b
}

// 匿名函数
var Adds = func(a, b int) int {
	return a + b
}

// 多个参数和多个返回值
func Swap(a, b int) (int, int) {
	return b, a
}

// 可变数量的参数
// more 对应 []int 切片类型
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

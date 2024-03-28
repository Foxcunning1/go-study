package main

import (
	"fmt"
	"sync"
)

var m sync.Map

func main() {
	m.Load("aa")
	var (
		//a []int               // nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
		//b = []int{}           // 空切片, 和 nil 不相等, 一般用来表示一个空的集合
		a = []int{1, 2, 3} // 有3个元素的切片, len和cap都为3
		//d = c[:2]             // 有2个元素的切片, len为2, cap为3
		//e = c[0:2:cap(c)]     // 有2个元素的切片, len为2, cap为3
		//f = c[:0]             // 有0个元素的切片, len为0, cap为3
		//g = make([]int, 3)    // 有3个元素的切片, len和cap都为3
		//h = make([]int, 2, 3) // 有2个元素的切片, len为2, cap为3
		//i = make([]int, 0, 3) // 有0个元素的切片, len为0, cap为3
	)
	a = append(a, 0) // 切片扩展1个空间
	i := 1
	x := 3
	copy(a[i+1:], a[i:]) // a[i:]向后移动1个位置
	a[i] = x             // 设置新添加的元素
	fmt.Printf("%v", a)
}

func TrimSpace(s []byte) []byte {
	b := s[:0]
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}

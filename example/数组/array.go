package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a)
	var b = [...]int{1, 2, 3}
	fmt.Println(b)
	var c = [...]int{2: 3, 1: 2} // 定义长度为3的int型数组, 元素为 0, 2, 3    : 前面是索引
	fmt.Println(c)
	var d = [...]int{1, 2, 4: 5, 6}
	fmt.Println(d)

	//数组的内存结构比较简单。比如下面是一个[4]int{2,3,5,7}数组值对应的内存结构：
}

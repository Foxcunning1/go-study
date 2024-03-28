package main

import "fmt"

func main() {
	var x uint8 = 255
	fmt.Println("x:", x)

	// 溢出后会回环到最小值
	x = x + 1
	fmt.Println("x after overflow:", x) // Output: x after overflow: 0

	// 更大的类型 uint16
	var y uint16 = 65535
	fmt.Println("y:", y)

	// 溢出后会回环到最小值
	y = y + 1
	fmt.Println("y after overflow:", y) // Output: y after overflow: 0
}

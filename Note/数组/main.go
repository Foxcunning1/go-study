package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%#v\n", []byte("Hello, 世界"))
	fmt.Println("\xe4\xb8\x96")                // 打印: 世
	fmt.Println("\xe7\x95\x8c")                // 打印: 界
	fmt.Println("\xe4\x00\x00\xe7\x95\x8cabc") // �界abc
	for i, c := range "\xe4\x00\x00\xe7\x95\x8cabc" {
		fmt.Println(i, c)
	}
}

package main

import "time"

//使用2个协程顺序打印一个字符串

func main() {
	ch := make(chan string, 1)

	go echo1(ch)
	<-ch
	go echo2(ch)
	<-ch

	time.Sleep(time.Second)
}

func echo1(ch chan string) {
	ch <- "hello"
	println("hello")
}

func echo2(ch chan string) {
	ch <- "world"
	println("world")
}

package main

import (
	"fmt"
	"sync"
)

//**问题描述**
//
//使用两个 `goroutine` 交替打印序列，一个 `goroutine` 打印数字， 另外一个 `goroutine` 打印字母， 最终效果如下：
//
//```bash
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
//```

func main() {
	numberGoroutine := make(chan int, 2)
	//aGoroutine := make(chan string,2)
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		i := 1
		for {
			fmt.Print(i)
			numberGoroutine <- i
			i++
			numberGoroutine <- i
			fmt.Print(i)
		}
	}()

	go func() {
		j := 'A'
		for {
			select {
			case <-numberGoroutine:
				fmt.Print(j)
			}
			j++

			if j == 'Z' {
				wg.Done()
				return
			}
		}

	}()

	wg.Wait()

}

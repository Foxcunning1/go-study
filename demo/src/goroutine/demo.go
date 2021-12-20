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
	numberGoroutine := make(chan bool)
	aGoroutine := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		i := 1
		for {
			select {
			case <-numberGoroutine:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				aGoroutine <- true
			}
		}
	}()

	go func(group *sync.WaitGroup) {
		j := 'A'
		for {

			select {
			case <-aGoroutine:
				fmt.Print(string(j))
				j++
				fmt.Print(string(j))
				j++
				numberGoroutine <- true
			}

			if j >= 'Z' {
				group.Done()
				return
			}
		}

	}(&wg)
	numberGoroutine <- true

	wg.Wait()

}

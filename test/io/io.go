package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 从标准输入读取
	data, err := ReadFrom(os.Stdin, 11)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

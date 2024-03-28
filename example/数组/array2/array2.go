package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func main() {
	var a = [...]int{1, 2, 3} // a 是一个数组
	var b = &a

	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])

	for k, v := range b {
		fmt.Println(k, v)
	}

	fmt.Println("--------------------")

	// 字符串数组
	//var s1 = [3]string{"a", "b", "c"}
	//var s2 = [...]string{"你好", "世界"}

	//结构体数组
	var line1 [2]image.Point
	fmt.Println(line1)
	var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	fmt.Println(line2)
	var line3 = [...]image.Point{{0, 0}, {1, 1}}
	fmt.Println(line3)

	fmt.Println("--------------------")

	// 图像解码器数组
	var decoder1 [2]func(io.Reader) (image.Image, error)

	fmt.Println(decoder1)
	var decoder2 = [...]func(io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}

	fmt.Println(decoder2)

	fmt.Println("--------------------")

	// 接口数组
	var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{123, "你好"}

	fmt.Println(unknown1)
	fmt.Println(unknown2)

}

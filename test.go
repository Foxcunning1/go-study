
package main

import (
	"errors"
	"fmt"
)


const SDY = 2
//如果需要，也可以明确指定常量的类型：
const Pi float32 = 3.1415926


//示例代码
var isActive bool  // 全局变量声明
var enabled, disabled = true, false  // 忽略类型的声明
func test() {
	var available bool  // 一般声明
	valid := false      // 简短声明
	available = true    // 赋值操作
	fmt.Printf("%t\n",available)
	fmt.Printf("%t\n",valid)
	fmt.Printf("c的值:%t,d的值:%t,e的值:%t\n",isActive,enabled,disabled)
}



func main() {

		//test()
	test3()
}

func test2()  {
	//定义一个名称为“variableName”，类型为"type"的变量
	var a string
	var b int8
	var c,d,e string = "2","3","4"
	/*
		定义三个变量，它们分别初始化为相应的值
		vname1为v1，vname2为v2，vname3为v3
		编译器会根据初始化的值自动推导出相应的类型
	*/
	vname1, vname2, vname3 := 6, 7, 8

	a = "ss"
	b = 14
	fmt.Printf("%s\n",a)
	fmt.Printf("%d\n",b)
	fmt.Printf("c的值:%s,d的值:%s,e的值:%s\n",c,d,e)
	fmt.Printf("c的值:%d,d的值:%d,e的值:%d\n",vname1,vname2,vname3)
	fmt.Printf("%d\n",SDY)
	fmt.Printf("%f\n",Pi)
}



//示例代码
var frenchHello string  // 声明变量为字符串的一般方法
var emptyString string = ""  // 声明了一个字符串变量，初始化为空字符串
func test1() {
	no, yes, maybe := "no", "yes", "maybe"  // 简短声明，同时声明多个变量
	japaneseHello := "Konichiwa"  // 同上
	frenchHello = "Bonjour"  // 常规赋值
	fmt.Printf("c的值:%s,d的值:%s,e的值:%s\n",no, yes, maybe)
	fmt.Printf("c的值:%s,d的值:%s,e的值:%s\n",japaneseHello, yes, maybe)




}

func test3()  {
	//var s string = "hello"
	//s[0] = 'c'

	s := "s我爱组合"
	//c := []byte(s)  // 将字符串 s 转换为 []byte 类型
	//c[0] = 'c'
	//s2 := string(c)  // 再转换回 string 类型
	//fmt.Printf("%s\n", s2)
	m := `hello
	world`


	fmt.Printf("s的值:%s\nm的值:%s\n",s[4:],m)


	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}

}


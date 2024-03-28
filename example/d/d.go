package d

import (
	"fmt"
)

func A() {
	fmt.Println("aaa")
	B()
}

func B() {
	fmt.Println("bbbbb")
	A()
}

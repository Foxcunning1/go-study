package main

import (
	"fmt"
	"math/rand"
)

//基于洗牌算法的负载均衡
var endpoints = []string{
	"100.69.62.1:3232",
	"100.69.62.32:3232",
	"100.69.62.42:3232",
	"100.69.62.81:3232",
	"100.69.62.11:3232",
	"100.69.62.113:3232",
	"100.69.62.101:3232",
}

func main() {
	var a = make(map[string]interface{})
	a["aa"] = "aa"

	err := request(a)
	if err != nil {
		panic(err)
	}
}

// 重点在这个 shuffle
func shuffle(slice []int) {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
}

func request(params map[string]interface{}) error {
	var indexes = []int{0, 1, 2, 3, 4, 5, 6}
	var err error

	shuffle(indexes)
	maxRetryTimes := 3

	idx := 0
	for i := 0; i < maxRetryTimes; i++ {
		fmt.Println("请求了", idx, indexes[idx])
		//err = apiRequest(params, indexes[idx])
		//if err == nil {
		//	break
		//}
		idx++
	}

	if err != nil {
		// logging
		return err
	}

	return nil
}

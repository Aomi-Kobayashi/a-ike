package main

import (
	"fmt"
	"math/rand"
)

// ランダム生成された1~10の整数で今日の運勢を占う
func main() {

	var num int = (rand.Intn(10))
	fmt.Println(num)

	if num == 7 || num == 9 {
		fmt.Println("Today is GodDay!")
	} else {
		fmt.Println("Well,it's Ok")
	}
}

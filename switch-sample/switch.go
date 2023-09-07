package main

import (
	"fmt"
	"strconv"
	"time"
)

// 現在時刻を取得し時刻に合わせて挨拶をする
func main() {
	// 「13:00」と出力したくstringをintに変換したがiが出力されない
	var s string = ":00"
	i, _ := strconv.Atoi(s)

	t := time.Now().Hour()

	fmt.Println(t + i)

	switch {
	case t < 12:
		fmt.Println("Good Morning")
	case t < 17:
		fmt.Println("Good Aftrnoon")
	case t == 22 || t == 23:
		fmt.Println("Good Night")
	default:
		fmt.Println("Good Evening")
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// type FPSGames struct {
// 	GameId         int      `json:"game-id"`
// 	GameName       string   `json:"game-name"`
// 	DevelopCompany string   `json:"company"`
// 	Oneteam        int      `json:"team-people"`
// 	Rank           []string `json:"rank"`
// }

// func readFile(file string)  {

// 	games, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		fmt.Println("file cannot read")
// 		return
// 	}

// 	fmt.Println(string(games))
// 	return
// }

func main() {

	// game.jsonを読み込み
	games, err := ioutil.ReadFile("game.json")
	if err != nil {
		fmt.Println("file cannot read", err)
		return
	}

	// sliceをstringに変換し文字列操作
	stringOperations := string(games)

	// 「Valorant」を検索
	fmt.Println(strings.Contains(stringOperations, "Valorant"))

	// 文字列の長さをカウント
	fmt.Println(len(stringOperations))

	// 9~411までを指定して表示
	fmt.Println(stringOperations[8:412])

}

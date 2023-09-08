package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 結合
	s := "VALORANT"
	s = s + "+ APEX"
	fmt.Println(s)

	// 大文字、小文字変換
	s = "ValorantApex"
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))

	// 部分取得
	s = "VALORANTApex"
	fmt.Println(s[0:9])
	fmt.Println(s[9:])
	fmt.Println(s[:9])
	fmt.Println(s[:]) // All

	// 両端トリム
	s = "     VALORANT    "
	fmt.Printf("[%s]\n", strings.TrimSpace(s))

	// トリム対象指定
	s = "     VALORANT    "
	fmt.Printf("[%s]\n", strings.Trim(s, " "))

	// 左、右トリム
	s = "     VALORANT     "
	fmt.Printf("[%s]\n", strings.TrimLeft(s, " "))

	s = "     VALORANT     "
	fmt.Printf("[%s]\n", strings.TrimRight(s, " "))

	// 指定文字が含まれているか
	s = "Apple Bear Cookie Dog Earth Fox Git"
	fmt.Println(strings.Contains(s, "Cookie"))
	fmt.Println(strings.Contains(s, "Happy"))

	// 指定文字列の出現位置
	s = "Apple Bear Cookie Dog Earth Fox Git"
	fmt.Println(strings.Index(s, "Cookie"))
	fmt.Println(strings.Index(s, "Happy"))

	// 指定文字列の最後の出現位置
	s = "Apple Bear Cookie Dog Earth Fox Git"
	fmt.Println(strings.LastIndex(s, "o"))

	// 先頭一致
	s = "Apple Bear Cookie Dog Earth Fox Git"
	fmt.Println(strings.HasPrefix(s, "Apple"))

	// 後方一致
	s = "Apple Bear Cookie Dog Earth Fox Git"
	fmt.Println(strings.HasSuffix(s, "Git"))

	// 文字列変換
	s = "bookbookbookbook"
	fmt.Println(strings.Replace(s, "oo", "uu", 3))  // 3個
	fmt.Println(strings.Replace(s, "oo", "ee", -1)) // All

	// 繰り返し
	s = "apex"
	fmt.Println(strings.Repeat(s, 4)) // 4回繰り返し

	// タイトルケースへの変換
	s = "apple bear cookie dog earth fox git"
	fmt.Println(strings.Title(s))

	// 分割
	s = "Apple Bear Cookie Dog Earth Fox Git"
	slice := strings.Split(s, " ") // space区切りで分割
	for _, str := range slice {    // sliceの数だけforで回す
		fmt.Printf("[%s]", str)
	}
	fmt.Println("")

	// 結合2
	slice = []string{"Apple", "Bear", "Cookie", "Dog", "Earth", "Fox", "Git"}
	fmt.Println(strings.Join(slice, ","))

	// 文字列長カウント
	s = "Apple Bear Cookie Dog Earth Fox Git"
	fmt.Println(len(s))

	// 文字列長カウント(マルチバイト)
	s = "Java Go言語"
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))

	// 複数行文字列
	s = `
	This is a sample.
	golang multiple lines.
	`
	fmt.Println(s)

	// 文字列→浮動小数点変換
	s = "990.339"
	sf, _ := strconv.ParseFloat(s, 64) // sfに代入
	sf = sf + 0.1
	fmt.Println(sf)
}

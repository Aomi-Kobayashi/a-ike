package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
)

func endcoding(src string) bool {
	// 変換したい文字列をキャストして変数に格納
	data := []byte(src)

	// エンコード
	encoded := base64.StdEncoding.EncodeToString(data)

	fmt.Println(encoded)

	// デコード
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Fatal(err)
		return false
	}
	// byteからstringにキャストして表示
	fmt.Println(string(decoded))
	return true
}

func main() {
	src := "Hello World"
	endecordedResult := endcoding(src)

	if endecordedResult {
		fmt.Println("Succsess")
		fmt.Println()
	} else {
		fmt.Println("Could not decode.")
	}

	// パーセントエンコード
	urlData := "インスタグラム"
	urlData = url.QueryEscape(urlData)

	fmt.Println(urlData) // %E3%82%A4%E3%83%B3%E3%82%B9%E3%82%BF%E3%82%B0%E3%83%A9%E3%83%A0

	// パーセントデコード
	urlData, _ = url.QueryUnescape(urlData)

	fmt.Println(urlData) // インスタグラム

	// URLエンコード
	urlencoded := base64.URLEncoding.EncodeToString([]byte(urlData))
	fmt.Println(urlencoded) // 44Kk44Oz44K544K_44Kw44Op44Og

	// URLデコード
	urldecoded, _ := base64.URLEncoding.DecodeString(urlencoded)
	fmt.Println(string(urldecoded)) // インスタグラム
}

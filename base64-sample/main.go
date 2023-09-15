package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

// Base64のエンコード、デコードを行う
func endcoding(src string) bool {
	// 変換したい文字列をキャストして変数に格納
	data := []byte(src)

	// エンコード
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	// デコード
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Decode is failed", err)
		return false
	}
	// byteからstringにキャストして表示
	fmt.Println(string(decoded))
	return true
}

// パーセントのエンコード、デコードを行う
func parcentEndecode(urlData string) bool {

	// エンコード
	urlData = url.QueryEscape(urlData)
	fmt.Println(urlData) // %E3%82%A4%E3%83%B3%E3%82%B9%E3%82%BF%E3%82%B0%E3%83%A9%E3%83%A0

	// デコード
	urlData, err := url.QueryUnescape(urlData)
	if err != nil {
		fmt.Println("Decode is failed", err)
		return false
	}
	fmt.Println(urlData) // インスタグラム
	return true
}

// URLのエンコード、デコードを行う
func urlEndecode(urlData string) bool {

	// エンコード
	urlencoded := base64.URLEncoding.EncodeToString([]byte(urlData))
	fmt.Println(urlencoded) // 44Kk44Oz44K544K_44Kw44Op44Og

	// デコード
	urldecoded, err := base64.URLEncoding.DecodeString(urlencoded)
	if err != nil {
		fmt.Println("Decode is failed", err)
		return false
	}
	fmt.Println(string(urldecoded)) // インスタグラム
	return true
}

func main() {

	// Base64の処理結果
	src := "Hello World" // データ代入
	endecordedResult := endcoding(src)

	if !endecordedResult {
		fmt.Println("Could not decode")
		return
	}
	fmt.Println("Base64 is Succsess")

	// パーセントの処理結果
	urlData := "インスタグラム" // データ代入
	parcentEndecodeResult := parcentEndecode(urlData)

	if !parcentEndecodeResult {
		fmt.Println("Processing failed")
		return
	}
	fmt.Println("Parcent is Succsess")

	// URLの処理結果
	urlEndecodeResult := urlEndecode(urlData)

	if !urlEndecodeResult {
		fmt.Println("Processing failed")
		return
	}
	fmt.Println("Url is Succsess")
}

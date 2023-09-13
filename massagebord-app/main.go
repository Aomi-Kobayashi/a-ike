package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"
)

const logFile = "logs.json"

type Log struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Kinds string `json:kinds"`
	Body  string `json:"body"`
	PTime int64  `json:ptime"`
}

func readLogs() []Log {
	text, err := ioutil.ReadFile(logFile)
	if err != nil {
		return make([]Log, 0)
	}

	var logs []Log
	json.Unmarshal([]byte(text))
}

// リクエストを処理する関数
// func menuHandler(w http.ResponseWriter, r *http.Request) {
// 	html, err := template.ParseFiles("menu.html")
// 	if err != nil {
// 		log.Fatal(err)

// 	}
// 	if err := html.Execute(w, nil); err != nil {
// 		log.Fatal(err)

// 	}
// }

func main() {
	// リクエストを処理する関数を登録
	http.HandleFunc("/", menuHandler)

	// Webサーバの設定
	server := http.Server{
		// ホスト名とポート番号
		Addr:    ":8080",
		Handler: nil,
		// リクエスト読み取りタイムアウト
		ReadTimeout: 30 * time.Second,
		// レスポンス書き込みのタイムアウト
		WriteTimeout: 60 * time.Second,
		// リクエストヘッダの最大バイト長
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Server Start Up........")
	// Webサーバの起動
	server.ListenAndServe()
}

// リクエストを受け付ける処理
func menuHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("menu.html")
	if err != nil {
		panic("fact: html file not found")
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

// // リクエストを受け付ける処理（ハンドラ）
// func defaultRoute(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "This is default route.")
// }

// // リクエストを受け付ける処理（ハンドラ）
// func route1(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "This is /route1.")
// }

// // リクエストを受け付ける処理（ハンドラ）
// func route2(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "This is /route2.")
// }

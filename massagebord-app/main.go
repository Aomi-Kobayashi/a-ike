package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"time"
)

// 書き込みログをjsonで保存
const logFile = "logs.json"

// ログの構造体を定義
type Log struct {
	ID        int    `json:"id"`       // 自動採番
	Name      string `json:"name"`     // 投稿者名
	GameTitle string `json:gametitle"` // ゲームの種類
	Body      string `json:"body"`     // 本文
	PTime     int64  `json:ptime"`     // 投稿時間
}

func main() {
	println("server - http://localhost:8888")

	http.HandleFunc("/", showHandler)
	http.HandleFunc("/write", writeHandler)

	http.ListenAndServe(":8888", nil)
}

func showHandler(w http.ResponseWriter, r *http.Request) {

	htmlLog := ""
	logs := loadLogs()
	for _, i := range logs {
		htmlLog += fmt.Sprintf(
			"<p>(%d) <span>%s</span>: %s --- %s</p>",
			i.ID,
			html.EscapeString(i.Name),
			html.EscapeString(i.Body),
			time.Unix(i.PTime, 0).Format("2006/1/2 15:04"))
	}

	htmlBody := "<html><head><style>" +
		"p { border: 1px solid silver; padding: 1em;} " +
		"span { background-color: #eef; } " +
		"</style></head><body><h1>掲示板</h1>" +
		getForm() + htmlLog + "</body></html>"
	w.Write([]byte(htmlBody))
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var log Log
	log.Name = r.Form["name"][0]
	log.Body = r.Form["body"][0]
	log.GameTitle = r.Form["gametitle"][0]
	if log.Name == "" {
		log.Name = "名無し"
	}
	logs := loadLogs()
	log.ID = len(logs) + 1
	log.PTime = time.Now().Unix()
	logs = append(logs, log)
	saveLogs(logs)
	http.Redirect(w, r, "/", 302)

}
func getForm() string {
	return "<div><form action='/write' method='POST'>" +
		"名前: <input type='text' name='name'><br>" +
		"ゲームタイトル: <input type='text' name='gametitle'><br>" +
		"本文: <input type='text' name='body' style='width:30em;'>" +
		"<br><input type='submit' value='書込'>" +
		"</form></div><hr>"
}

// ログファイルの読み込み
func loadLogs() []Log {
	// ファイルを読む
	text, err := ioutil.ReadFile(logFile)
	if err != nil {
		fmt.Println("Could not read file", err)
		// 初期値0のLog型配列を作成し返す
		return make([]Log, 0)
	}

	var logs []Log
	json.Unmarshal([]byte(text), &logs)
	return logs
}

func saveLogs(logs []Log) {
	bytes, _ := json.Marshal(logs)
	ioutil.WriteFile(logFile, bytes, 0644)
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

// func main() {
// 	// リクエストを処理する関数を登録
// 	http.HandleFunc("/", menuHandler)

// 	// Webサーバの設定
// 	server := http.Server{
// 		// ホスト名とポート番号
// 		Addr:    ":8080",
// 		Handler: nil,
// 		// リクエスト読み取りタイムアウト
// 		ReadTimeout: 30 * time.Second,
// 		// レスポンス書き込みのタイムアウト
// 		WriteTimeout: 60 * time.Second,
// 		// リクエストヘッダの最大バイト長
// 		MaxHeaderBytes: 1 << 20,
// 	}

// 	fmt.Println("Server Start Up........")
// 	// Webサーバの起動
// 	server.ListenAndServe()
// }

// // リクエストを受け付ける処理
// func menuHandler(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("menu.html")
// 	if err != nil {
// 		panic("fact: html file not found")
// 	}
// 	if err := t.Execute(w, nil); err != nil {
// 		panic(err.Error())
// 	}
// }

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

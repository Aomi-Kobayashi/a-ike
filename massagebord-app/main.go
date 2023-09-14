package main

import (
	"bytes"
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

// サーバを起動する
func main() {
	println("server - http://localhost:8888") // 案内
	// パスとハンドラ関数を結びつける
	http.HandleFunc("/", showHandler)
	// パスとハンドラ関数を結びつける
	http.HandleFunc("/write", writeHandler)
	// サーバ起動 Getで何も渡さないのでnil
	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println("Server startup failed", err)
	}
	fmt.Println("Succese")
}

// 書き込んだものを画面に表示するハンドラ作成
func showHandler(writeRes http.ResponseWriter, req *http.Request) {
	// ファイルからログを読み込んで、各ログをhtmlに格納
	htmlLog := ""
	logs := loadLogs()
	// jsonの中身をforで回してiに格納
	for _, i := range logs {
		htmlLog += fmt.Sprintf(
			"<p>(%d) <span>%s</span>: %s --- %s</p>",
			i.ID,
			html.EscapeString(i.Name), // htmlにする文字列をエスケープ
			html.EscapeString(i.Body),
			time.Unix(i.PTime, 0).Format("2006/1/2 15:04"))
	}

	// html全体を出力
	htmlBody := "<html><head><style>" +
		"body { position: relative; width: auto; height: auto; text-align: center;}" +
		"form {position: relative; z-index: 1; background: #FFFFFF; max-width: 360px; margin: 0 auto 100px; padding: 45px; text-align: center;}" +
		"p { border: 1px solid silver; padding: 1em;} " +
		"span { background-color: #eef; } " +
		"</style></head><body><h1>ゲームフレンド募集掲示板</h1>" +
		getForm() + htmlLog + "</body></html>"
	// htmlをキャストして書き込む
	writeRes.Write([]byte(htmlBody))
}

// フォームから送信された内容を書き込み
func writeHandler(writeRes http.ResponseWriter, req *http.Request) {
	req.ParseForm() // フォームを解析
	var log Log
	log.Name = req.Form["name"][0]
	log.Body = req.Form["body"][0]
	log.GameTitle = req.Form["gametitle"][0]
	// Nameが空白の場合、名無しとして書き込み
	if log.Name == "" {
		log.Name = "名無し"
	}
	logs := loadLogs()     // 既存データを読み出し
	log.ID = len(logs) + 1 // 0からなので+1して1から
	log.PTime = time.Now().Unix()
	logs = append(logs, log)               // 既存データに新規書き込みを追記
	saveLogs(logs)                         //保存
	http.Redirect(writeRes, req, "/", 302) // ログ表示を行うルートページへリダイレクト
}

// 画面上部の書き込みフォームを返す/writeに向けてPostで送信
func getForm() string {
	return "<div><form action='/write' method='POST'>" +
		"名前 <input type='text' name='name' style='width:25em;'><br>" +
		"ゲームタイトル <input type='text' name='gametitle' style='width:25em;'><br>" +
		"本文 <input type='text' name='body' style='width:25em; height:10em;'>" +
		"<br><input type='submit' id='write' value='書込'>" +
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
	// jsonをパース
	var logs []Log
	if err := json.Unmarshal([]byte(text), &logs); err != nil {
		fmt.Println("Parsing failed", err)
		// 初期値0のLog型配列を作成し返す
		return make([]Log, 0)
	}

	return logs // 成功したらlogsに返す
}

// 書き込みログを保存する
func saveLogs(logs []Log) {
	bytes2, _ := json.Marshal(logs)
	out := new(bytes.Buffer)             // バッファ作成
	json.Indent(out, bytes2, "", "    ") // JSONテキストの成形
	ioutil.WriteFile(logFile, []byte(out.String()), 0644)
}

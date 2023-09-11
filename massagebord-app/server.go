package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func menuHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("menu.html")
	if err != nil {
		log.Fatal(err)

	}
	if err := html.Execute(w, nil); err != nil {
		log.Fatal(err)

	}
}

func main() {
	http.HandleFunc("/menu", menuHandler)
	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

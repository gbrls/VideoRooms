package main

// Gabriel Schneider - 2019

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	timeValue float64
	changed   chan int
	conn      int
)

func isAdmin(r *http.Request) bool {
	//TODO: implement this with hash maps.

	chars := []rune(r.RemoteAddr)
	if chars[0] == '[' {

		fmt.Println(string(chars), "is admin")

		return true
	}

	fmt.Println(string(chars), "is NOT admin")

	return false
}

func main() {
	data, err := ioutil.ReadFile("./html/index.gtpl")
	if err != nil {
		log.Fatal(err)
	}

	changed = make(chan int)
	conn = 0

	videoTmp := template.Must(template.New("").Parse(string(data)))
	ws := websocket.Upgrader{}

	log.Println("Starting server...")

	http.HandleFunc("/video", videoPlayer(videoTmp))
	http.HandleFunc("/ws", wsHandler(&ws))
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/upload", upload)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

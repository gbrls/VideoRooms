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

func loadTemplate(file string) *template.Template {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return template.Must(template.New("").Parse(string(data)))
}

func main() {

	changed = make(chan int)
	conn = 0
	mainApp := newApp()

	videoTmp := loadTemplate("./html/index.gtpl")
	homeTmp := loadTemplate("./html/home.gtpl")
	formTmp := loadTemplate("./html/form.gtpl")

	ws := websocket.Upgrader{}

	log.Println("Starting server...")

	http.HandleFunc("/r/", videoPlayer(videoTmp))
	http.HandleFunc("/ws", wsHandler(&ws, mainApp))
	http.HandleFunc("/form", formHandler(formTmp))
	http.HandleFunc("/upload", upload(mainApp))
	http.HandleFunc("/", homeHandler(homeTmp))

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

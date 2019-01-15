package main

// Gabriel Schneider - 2019

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var (
	timeValue float64
	changed   chan int
	conn      int
)

func isAdmin(r *http.Request) bool {
	//TODO: implement this from room struct.

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

	cfgFile := flag.String("cfg", "./cfg.yaml", "Path to the configuration file")

	flag.Parse()

	f, err := os.Open(*cfgFile)
	if err != nil {
		log.Fatalf("didnt find cfg file (%s)", err)
	}

	mainApp := newApp(f)

	changed = make(chan int)
	conn = 0

	videoTmp := loadTemplate("./html/index.gtpl")
	homeTmp := loadTemplate("./html/home.gtpl")
	formTmp := loadTemplate("./html/form.gtpl")

	ws := websocket.Upgrader{}

	log.Println("Starting server...")

	http.HandleFunc("/r/", videoPlayer(videoTmp, mainApp))
	http.HandleFunc("/ws/", wsHandler(&ws, mainApp))
	http.HandleFunc("/form", formHandler(formTmp, mainApp))
	http.HandleFunc("/upload", upload(mainApp))
	http.HandleFunc("/", homeHandler(homeTmp, mainApp))

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":"+mainApp.Port, nil))
}

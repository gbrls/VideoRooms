package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var (
	timeValue float64
)

func home(tmp *template.Template) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		tmp.Execute(w, "ws://"+r.Host+"/ws")
	}
}

func wsHandler(ws *websocket.Upgrader) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		c, err := ws.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}

		defer c.Close()

		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				//log.Println(err)
				break
			}
			timeValue, _ = strconv.ParseFloat(string(message), 64)

			c.WriteMessage(mt, []byte(fmt.Sprintf("%0.5f", timeValue)))

			fmt.Printf("sending %0.4f\n", timeValue)
			fmt.Printf("recived %v %v\n", mt, string(message))

		}

		fmt.Println("Connection closed.")

	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This is a form")

}

func main() {
	data, err := ioutil.ReadFile("./html/index.gtpl")
	if err != nil {
		log.Fatal(err)
	}

	homeTmp := template.Must(template.New("").Parse(string(data)))
	ws := websocket.Upgrader{}

	log.Println("Starting server...")

	http.HandleFunc("/", home(homeTmp))
	http.HandleFunc("/ws", wsHandler(&ws))
	http.HandleFunc("/form", formHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8082", nil))
}

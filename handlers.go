package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

func videoPlayer(tmp *template.Template) func(http.ResponseWriter, *http.Request) {

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

		log.Printf("Connected to (%v, %v) \n", r.RemoteAddr, r.UserAgent())

		conn++

		go func() {
			for {
				n := <-changed

				if n > 1 {
					fmt.Printf("Sending new times to %v/%v\n", conn-(n-1), conn)

					changed <- (n - 1)
				}

				if !isAdmin(r) {
					c.WriteMessage(1, []byte(fmt.Sprintf("%0.5f", timeValue)))

				}
			}
		}()

		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				//log.Println(err)
				break
			}

			if isAdmin(r) {
				timeValue, _ = strconv.ParseFloat(string(message), 64)
				c.WriteMessage(mt, []byte(fmt.Sprintf("%0.5f", timeValue)))

				changed <- conn

			} else {

				c.WriteMessage(mt, []byte(fmt.Sprintf("%0.5f", timeValue+0.5)))
			}

			//log.Printf("Sending (%0.2f) to %v\n", timeValue, r.RemoteAddr)
			//log.Printf("Recived (type:%v message:%v) from %v\n", mt, string(message), r.RemoteAddr)

		}

		log.Printf("Connection with (%v) closed.\n", r.RemoteAddr)
		conn--

	}

}

func formHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This is a form")

}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "you shoudn't be here ;)")

	}
}

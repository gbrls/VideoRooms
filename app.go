package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

const maxrooms = 10

//Application stores data about the app
type Application struct {
	rooms map[int]*room

	Port      string `yaml:"port"`
	Host      string `yaml:"host"`
	Localhost bool   `yaml:"localhost"`
}

func newApp(cfg io.Reader) *Application {
	data, err := ioutil.ReadAll(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	var app Application
	err = yaml.Unmarshal(data, &app)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v decoded\n", app)

	app.rooms = make(map[int]*room)

	if app.Localhost {
		app.rooms[0] = newRoom("", "filme.mp4")
		fmt.Println("Creating local room")
	}

	return &app
}

func (app *Application) addRoom(admin string, video string) {
	var i int

	for i = 0; i < maxrooms; i++ {
		if app.rooms[i] == nil {
			break
		}
	}

	app.rooms[i] = newRoom(admin, video)
}

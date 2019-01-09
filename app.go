package main

type application struct {
	rooms map[int]*room
}

func newApp() *application {
	return &application{}
}

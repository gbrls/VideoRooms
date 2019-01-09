package main

type room struct {
	admin          string
	usersConnected int
}

func newRoom() *room {
	return &room{}
}

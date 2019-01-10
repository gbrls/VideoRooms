package main

type room struct {
	admin          string
	usersConnected int
	videoName      string
}

func newRoom(admin string, video string) *room {
	return &room{
		admin:          admin,
		usersConnected: 0,
		videoName:      video,
	}
}

package main

const maxrooms = 10

type application struct {
	rooms map[int]*room
}

func newApp() *application {
	app := &application{}

	app.rooms = make(map[int]*room)

	app.rooms[0] = newRoom("", "filme.mp4")

	return app
}

func (app *application) addRoom(admin string, video string) {
	var i int

	for i = 0; i < maxrooms; i++ {
		if app.rooms[i] == nil {
			break
		}
	}

	app.rooms[i] = newRoom(admin, video)
}

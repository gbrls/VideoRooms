package main

import (
	"io"
	"net/http"
	"os"
	"regexp"
)

func saveFileToDisk(r *http.Request, filename string, pattern string) (string, error) {

	file, handler, err := r.FormFile(filename)

	defer file.Close()

	res, err := regexp.MatchString(pattern, handler.Filename)
	if !res {
		//TODO: fix this

		//fmt.Println("is not a video")
		_, _ = err, res
		//return "", fmt.Errorf("%s is not a video", handler.Filename)
	}

	f, err := os.OpenFile("./static/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	defer f.Close()
	io.Copy(f, file)

	return handler.Filename, err
}

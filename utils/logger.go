package utils

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	file, err := os.OpenFile(
		"logs/app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	Logger = log.New(
		io.MultiWriter(os.Stdout, file),
		"[GO-DOCKER-API] ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}

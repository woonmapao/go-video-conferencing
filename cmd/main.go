package main

import (
	"go-video-conferencing/internal/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

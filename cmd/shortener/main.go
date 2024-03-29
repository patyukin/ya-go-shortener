package main

import (
	"log"

	"github.com/patyukin/ya-go-shortener/internal/app/server"
	"github.com/patyukin/ya-go-shortener/internal/app/shortener"
	"github.com/patyukin/ya-go-shortener/internal/app/storage"
)

func main() {
	s := shortener.New()
	st := storage.New()
	err := server.Serve("127.0.0.1:8080", s, st)
	if err != nil {
		log.Fatal(err)
	}
}

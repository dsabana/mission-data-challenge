package main

import (
	"fmt"
	"log"
	http2 "mission-data-challenge/internal/http"
	"mission-data-challenge/internal/service"
	"mission-data-challenge/internal/storage"
	"net/http"
)

func main() {
	storage, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	service := service.NewService(storage)

	router := http2.SetupRouter(service)

	fmt.Printf("Serving transactions on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}

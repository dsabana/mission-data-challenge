package main

import (
	"fmt"
	"log"
	"mission-data-challenge/internal"
	"net/http"
)

func main() {
	storage, err := internal.NewStorage()
	if err != nil {
		panic(err)
	}

	service := internal.NewService(storage)

	router := internal.SetupRouter(service)

	fmt.Printf("Serving transactions on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}

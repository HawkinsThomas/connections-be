package main

import (
	"connections/router"
	"log"
	"net/http"
)

func main() {
	r := router.InitializeRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

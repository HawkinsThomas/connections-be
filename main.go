
package main

import (
	"net/http"
	"log"
	"connections/router"
)

func main() {
	r := router.InitializeRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
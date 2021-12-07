package main

import (
	"fmt"
	"log"
	"net/http"
	"solaris/router"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("/assets/"))
	http.Handle("/", r)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}

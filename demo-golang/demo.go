package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

/*
creates a basic web server which listens on port 18010:
*/
func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "World"
	}
	fmt.Fprintf(w, "startkit-go: Hello %s!\n", target)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\":\"UP\"}")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "18010"
	}

	log.Printf("Golang-Hello world sample started. listen on: %s \n", port)

	http.HandleFunc("/", handler)
	http.HandleFunc("/health", health)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

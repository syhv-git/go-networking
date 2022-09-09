package main

import (
	"go-networking/frontend"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", frontend.Build)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

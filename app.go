package main

import (
	"go-networking/frontend"
	"log"
	"net/http"
)

func main() {
	mux := frontend.Routes()

	err := http.ListenAndServe(":9001", mux)
	log.Fatal(err)
}

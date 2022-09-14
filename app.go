package main

import (
	"go-networking/frontend"
	"log"
	"net/http"
)

func main() {
	mux := frontend.Routes()

	log.Fatal(http.ListenAndServe(":9001", mux))
}

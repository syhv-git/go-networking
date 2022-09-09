package frontend

import (
	"go-networking/frontend/pages"
	"net/http"
)

func Routes() *http.ServeMux {
	m := http.NewServeMux()

	m.HandleFunc("/", pages.Build)

	m.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	return m
}

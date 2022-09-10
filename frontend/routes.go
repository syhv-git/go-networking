package frontend

import (
	"go-networking/frontend/pages"
	"net/http"
)

func Routes() *http.ServeMux {
	m := http.NewServeMux()

	a := getAssets()
	m.Handle("/assets/", a)
	m.HandleFunc("/", pages.Build)

	return m
}

func getAssets() http.Handler {
	d := http.Dir("../assets/")
	return http.StripPrefix("/assets/", http.FileServer(d))
}

package frontend

import (
	"github.com/gorilla/mux"
	"go-networking/frontend/pages"
	"net/http"
)

func Routes(dir string) *mux.Router {
	a := getAssets(dir)
	m := mux.NewRouter()
	m.PathPrefix("/assets/").Handler(a).Methods("GET")

	m.HandleFunc("/", pages.Build).Methods("GET")
	m.HandleFunc("/search-results", pages.SearchResults).Methods("GET")

	return m
}

func getAssets(dir string) http.Handler {
	return http.StripPrefix("/assets/", http.FileServer(http.Dir(dir)))
}

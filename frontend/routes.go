package frontend

import (
	"go-networking/frontend/pages"
	"net/http"
)

func Routes() *http.ServeMux {
	m := http.NewServeMux()

	m.HandleFunc("/", pages.Build)

	return m
}

package pages

import (
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func Build(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	if err := tpl.Execute(w, nil); err != nil {
		panic(err)
	}
}

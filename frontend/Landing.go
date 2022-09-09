package frontend

import "net/http"

func Build(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("<h1>Hello World!</h1>"))
	if err != nil {
		panic(err)
	}
}

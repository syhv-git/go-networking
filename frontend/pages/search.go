package pages

import (
	"fmt"
	"net/http"
	"net/url"
)

func BuildSearch(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := u.Query()
	query := param.Get("q")
	page := param.Get("page")
	if page == "" {
		page = "1"
	}
	fmt.Println(query)
}

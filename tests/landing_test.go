package tests

import (
	"go-networking/frontend/pages"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestLandingPageContent(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:9001", nil) //update url when deployed
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	h := http.HandlerFunc(pages.Build)
	h.ServeHTTP(rec, req)

	e, err := os.ReadFile("index.html")
	if err != nil {
		t.Fatal(err)
	}
	a := rec.Body.String()
	if a != string(e) {
		t.Errorf("** Failed: Unexpected response body: got %v want %v", a, e)
	}
}

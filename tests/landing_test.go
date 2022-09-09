package tests

import (
	"go-networking/frontend"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDomain(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:9001", nil) //update url when deployed
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	build := http.HandlerFunc(frontend.Build)
	build.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("** Failed: Response from server was %v; expected %v", status, http.StatusOK)
	}
}

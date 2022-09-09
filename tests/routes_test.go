package tests

import (
	"bytes"
	"go-networking/frontend"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	r := frontend.Routes()
	mock := httptest.NewServer(r)

	resp, err := http.Get(mock.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("** Failed: Response from server was %v; expected %v", resp.StatusCode, http.StatusOK)
	}

	resp, err = http.Post(mock.URL+"/", "text/plain", bytes.NewReader([]byte("Hello Error!")))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("** Failed: Response from server was %v; expected %v", resp.StatusCode, http.StatusMethodNotAllowed)
	}

}

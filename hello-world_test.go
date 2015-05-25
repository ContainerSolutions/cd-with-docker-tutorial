package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	homeHandle := homeHandler()
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	homeHandle.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}

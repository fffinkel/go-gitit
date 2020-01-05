package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot200(t *testing.T) {
	registerHandlers()

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(rr, req)

	res := rr.Result()

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.StatusCode)
		t.Errorf("no dice")
	}
}

package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIpHandlerStatusCode(t *testing.T) {
	want := 200
	handler := IpHandler
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ip", nil)
	if err != nil {
		t.Fatal("Error when creating the /ip GET request")
	}

	handler(recorder, req)

	resp := recorder.Result()

	if want != resp.StatusCode {
		t.Fatal("Bad status code")
	}
}

func TestIpHandlerHeader(t *testing.T) {
	want := "application/json"
	handler := IpHandler
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ip", nil)
	if err != nil {
		t.Fatal("Error when creating the /ip GET request")
	}

	handler(recorder, req)

	resp := recorder.Result()

	if want != resp.Header.Get("Content-Type") {
		t.Fatal("Bad Content-Type")
	}
}

func TestIpHandlerBody(t *testing.T) {
	handler := IpHandler
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ip", nil)
	if err != nil {
		t.Fatal("Error when creating the /ip GET request")
	}

	want := "{\"ip\":\"" + req.RemoteAddr + "\"}"

	handler(recorder, req)

	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)

	if want != string(body) {
		t.Fatalf("Bad Body, got: %s want: %s", string(body), want)
	}
}

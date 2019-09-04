package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type fnRoute func(w http.ResponseWriter, r *http.Request)

func TestShouldReturnHostnameWhenSet(t *testing.T) {
	//given
	setHostname("TEST")
	//when
	_, err := GetHostname()
	//then
	if err != nil {
		t.Errorf("HOSTNAME SHOULD NOT be Empty")
	}
}

func TestShouldReturnErroWhenHostnameIsEmpty(t *testing.T) {
	//given
	setHostname("")
	//when
	_, err := GetHostname()
	//then
	if err == nil {
		t.Errorf("SHOULD return error when HOSTNAME is empty! %s", err.Error())
	}
}

func TestShouldGet200WhenRequestGetHostnameHandleWhenSet(t *testing.T) {
	//given
	setHostname("TEST")

	//when
	rr, _ := testRequest(GetHostNameHandler, "GET", "/")

	//then
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//then
	actual := rr.Body.Bytes()
	expected := []byte("Hello World! TEST")
	if bytes.Equal(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestShouldGet500WhentGetHostnameHandleWhenNotSet(t *testing.T) {
	//given
	setHostname("")

	//when
	rr, _ := testRequest(GetHostNameHandler, "GET", "/")

	//then
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	//then
	actual := rr.Body.Bytes()
	expected := []byte("Hello World! TEST")
	if bytes.Equal(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func testRequest(fn fnRoute, verb string, endpoint string) (*httptest.ResponseRecorder, *http.Request) {
	req, err := http.NewRequest(verb, endpoint, nil)
	if err != nil {
		panic(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fn)
	handler.ServeHTTP(rr, req)

	return rr, req
}

func setHostname(hostname string) {
	if err := os.Setenv("HOSTNAME", hostname); err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	HelloWorld(w, req)
	
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	got := string(data)
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestHelloUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/Katy", nil)
	w := httptest.NewRecorder()

	// Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"name": "Katy",
	}

	req = mux.SetURLVars(req, vars)

	HelloUser(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	got := string(data)
	want := "Hello, Katy!"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

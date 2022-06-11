package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func buildRequest(target string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, target, nil)
	w := httptest.NewRecorder()

	return req, w
}

func getOutput(w *httptest.ResponseRecorder) (string, error) {
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	return string(data), err
}

func TestSayHello(t *testing.T) {
	t.Run("It says hello world when no user is specified", func(t *testing.T) {
		req, w := buildRequest("/")

		SayHello(w, req)

		output, err := getOutput(w)

		got := output
		want := "Hello, world!"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}

		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}
	})

	t.Run("It says the users name when a name is specified", func(t *testing.T) {
		req, w := buildRequest("/user/katy")

		// Hack to try to fake gorilla/mux vars
		vars := map[string]string{
			"name": "Katy",
		}

		req = mux.SetURLVars(req, vars)

		SayHello(w, req)

		output, err := getOutput(w)

		got := output
		want := "Hello, Katy!"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}

		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}
	})
}

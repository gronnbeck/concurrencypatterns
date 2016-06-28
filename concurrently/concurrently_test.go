package concurrently

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_naive_get_all_the_things(t *testing.T) {
	server := mockServer()
	results := GetAllTheThings(server.URL)

	containsOrFail("/facebook", results, t)
	containsOrFail("/github", results, t)
	containsOrFail("/twitter", results, t)
}

func mockServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)

		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		fmt.Fprintln(w, r.URL.String())
	}))
	return server
}

func containsOrFail(str string, in []string, t *testing.T) {
	if !contains(str, in) {
		t.Logf("Expected %v to be contained in %v", str, in)
		t.Fail()
	}
}

func contains(e string, s []string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

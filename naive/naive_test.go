package naive

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
	results := getAllTheThings(server.URL)

	expect("/facebook", results[0], t)
	expect("/github", results[1], t)
	expect("/twitter", results[2], t)
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

func expect(expected, actual string, t *testing.T) {
	if expected != actual {
		t.Logf("Expected %v but was %v", expected, actual)
		t.Fail()
	}
}

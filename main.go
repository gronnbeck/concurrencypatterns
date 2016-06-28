package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gronnbeck/concurrencypatterns/concurrently"
	"github.com/gronnbeck/concurrencypatterns/naive"
)

func main() {

	go startServer()

	stopNaive := measure()
	res1 := naive.GetAllTheThings("http://localhost:8080")
	elapsedNaive := stopNaive()

	fmt.Printf("Naive elapsed: %v\n", elapsedNaive)
	fmt.Println(res1)

	stopConc := measure()
	res2 := concurrently.GetAllTheThings("http://localhost:8080")
	elapsedConc := stopConc()

	fmt.Printf("Concurrently elapsed: %v\n", elapsedConc)
	fmt.Println(res2)

}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.String()))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func measure() func() time.Duration {
	now := time.Now()

	return func() time.Duration {
		return time.Since(now)
	}
}

package concurrently

import (
	"github.com/gronnbeck/concurrencypatterns/naive"
)

// GetAllTheThings implemented using go routines and channels
func GetAllTheThings(service string) []string {
	c := make(chan string)
	client := naive.Client{BaseURL: service}

	go func() {
		fbRes, _ := client.Get("facebook")
		c <- *fbRes
	}()

	go func() {
		ghRes, _ := client.Get("github")
		c <- *ghRes
	}()

	go func() {
		twRes, _ := client.Get("twitter")
		c <- *twRes
	}()

	results := []string{}
	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}

	return results
}

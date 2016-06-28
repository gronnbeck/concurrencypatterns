package naive

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetAllTheThings implemented naively
func GetAllTheThings(service string) []string {
	c := Client{BaseURL: service}

	results := []string{}

	fbRes, _ := c.Get("facebook")

	if fbRes != nil {
		results = append(results, *fbRes)
	}

	ghRes, _ := c.Get("github")

	if ghRes != nil {
		results = append(results, *ghRes)
	}

	twRes, _ := c.Get("twitter")

	if twRes != nil {
		results = append(results, *twRes)
	}

	return results
}

// Client is a naive client needed to run some experients
type Client struct {
	BaseURL string
}

// Get fetches body from URL and turns it into a string
func (c Client) Get(path string) (*string, error) {
	url := fmt.Sprintf("%v/%v", c.BaseURL, path)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	byt, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	result := strings.Trim(string(byt), "\n")

	return &result, nil
}

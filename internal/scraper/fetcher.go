package scraper

import (
	"fmt"
	"net/http"
	"time"
)

type Fetcher struct {
	client *http.Client
}


// * Constructor Function
func NewFetcher() *Fetcher {
	return &Fetcher{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// * The Fetch Function
func (f *Fetcher) Fetch(url string) ([]byte, string, error) {

	resp, err := f.client.Get(url)

	if err != nil {
		return nil, "", fmt.Errorf("Failed Request: %w")
	}

	defer resp.Body.Close()

	body := make([]byte, resp.ContentLength)

	_, err = resp.Body.Read(body)

	if err != nil {
		return nil, "", fmt.Errorf("failed reading body: %w", err)
	}

	return body, resp.Status, nil

}
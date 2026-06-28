package scraper

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Fetcher struct {
	client *http.Client
}

// Constructor Function
func NewFetcher() *Fetcher {
	return &Fetcher{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Fetch downloads the webpage
func (f *Fetcher) Fetch(url string) ([]byte, string, error) {

	resp, err := f.client.Get(url)

	if err != nil {
		return nil, "", fmt.Errorf("failed request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, "", fmt.Errorf(
			"failed reading body: %w",
			err,
		)
	}

	return body, resp.Status, nil
}

package scraper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestFetcherFetch(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusOK)

			w.Write([]byte(
				"<html><body>Hello Ripper</body></html>",
			))
		}),
	)

	defer server.Close()


	fetcher := NewFetcher()


	body, status, err := fetcher.Fetch(server.URL)


	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}


	if status != "200 OK" {
		t.Fatalf(
			"expected status 200 OK got %s",
			status,
		)
	}


	if string(body) != "<html><body>Hello Ripper</body></html>" {
		t.Fatalf(
			"unexpected body: %s",
			string(body),
		)
	}
}
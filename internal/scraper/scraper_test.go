package scraper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScraperScrape(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {

				w.Write([]byte(`
					<html>

					<head>
						<title>
							Test Page
						</title>
					</head>


					<body>

						<h1>
							Hello
						</h1>


						<a href="/about">
							About
						</a>


					</body>

					</html>
				`))
			},
		),
	)

	defer server.Close()

	scraper := NewScraper()

	page, err := scraper.Scrape(
		server.URL,
	)

	if err != nil {
		t.Fatal(err)
	}

	if page.Title != "Test Page" {
		t.Fatalf(
			"expected Test Page got %s",
			page.Title,
		)
	}

	if len(page.Links) != 1 {
		t.Fatalf(
			"expected 1 link got %d",
			len(page.Links),
		)
	}

	if len(page.Headings) != 1 {
		t.Fatalf(
			"expected 1 heading got %d",
			len(page.Headings),
		)
	}
}

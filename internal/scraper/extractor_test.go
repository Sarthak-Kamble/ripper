package scraper

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestExtractorExtract(t *testing.T) {

	html := `
	<html>
	<head>
		<title>
			Ripper Page
		</title>
	</head>

	<body>

	<h1>
		Introduction
	</h1>

	<h2>
		Architecture
	</h2>

	<a href="https://example.com">
		Example
	</a>


	<a href="/docs">
		Docs
	</a>

	</body>
	</html>
	`

	doc, err := goquery.NewDocumentFromReader(
		strings.NewReader(html),
	)

	if err != nil {
		t.Fatal(err)
	}

	extractor := NewExtractor()

	page := extractor.Extract(
		doc,
		"https://test.com",
		"200 OK",
	)

	if page.Title != "Ripper Page" {
		t.Fatalf(
			"expected title Ripper Page got %s",
			page.Title,
		)
	}

	if len(page.Links) != 2 {
		t.Fatalf(
			"expected 2 links got %d",
			len(page.Links),
		)
	}

	if len(page.Headings) != 2 {
		t.Fatalf(
			"expected 2 headings got %d",
			len(page.Headings),
		)
	}
}

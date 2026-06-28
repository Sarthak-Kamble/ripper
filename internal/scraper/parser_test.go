package scraper

import (
	"strings"
	"testing"
)

func TestParserParse(t *testing.T) {

	html := []byte(`
	<html>
		<head>
			<title>
				Ripper Test
			</title>
		</head>

		<body>
			<h1>Hello Ripper</h1>

			<a href="https://example.com">
				Example
			</a>

		</body>
	</html>
	`)

	parser := NewParser()

	doc, err := parser.Parse(html)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	// title := doc.Find("title").Text()

	title := strings.TrimSpace(doc.Find("title").Text())

	link, exists := doc.Find("a").Attr("href")

	if !exists {
		t.Fatal("expected link")
	}

	if link != "https://example.com" {
		t.Fatalf(
			"expected example link got %s",
			link,
		)
	}

	if title != "Ripper Test" {

		t.Fatalf(
			"expected title Ripper Test, got %s",
			title,
		)
	}
}

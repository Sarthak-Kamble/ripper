package extractor

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)


func TestLinkExtractor(t *testing.T){


	html := `
	<html>
	<body>

	<a href="/about">
	About
	</a>

	<a href="https://google.com">
	Google
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


	extractor := NewLinkExtractor(
		"https://example.com",
	)


	links := extractor.Extract(doc)



	if len(links) != 2 {
		t.Fatalf(
			"expected 2 links got %d",
			len(links),
		)
	}



	if links[0].Type != "internal" {
		t.Fatal("expected internal link")
	}



	if links[1].Type != "external" {
		t.Fatal("expected external link")
	}

}
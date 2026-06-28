package scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(body []byte) (*goquery.Document, error) {

	doc, err := goquery.NewDocumentFromReader(
		strings.NewReader(string(body)),
	)

	if err != nil {
		return nil, err
	}

	return doc, nil
}

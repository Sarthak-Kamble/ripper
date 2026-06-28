package scraper

import (
	"github.com/Sarthak-Kamble/ripper/pkg/models"
)

type Scraper struct {
	fetcher   *Fetcher
	parser    *Parser
	extractor *Extractor
}

func NewScraper() *Scraper {

	return &Scraper{
		fetcher:   NewFetcher(),
		parser:    NewParser(),
		extractor: NewExtractor(),
	}
}

func (s *Scraper) Scrape(url string) (models.Page, error) {

	body, status, err := s.fetcher.Fetch(url)

	if err != nil {
		return models.Page{}, err
	}

	doc, err := s.parser.Parse(body)

	if err != nil {
		return models.Page{}, err
	}

	page := s.extractor.Extract(
		doc,
		url,
		status,
	)

	return page, nil
}

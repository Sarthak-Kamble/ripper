package scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Sarthak-Kamble/ripper/pkg/models"
)

type Extractor struct{}

func NewExtractor() *Extractor {
	return &Extractor{}
}

func (e *Extractor) Extract(
	doc *goquery.Document,
	url string,
	status string,
) models.Page {

	page := models.Page{
		URL:    url,
		Status: status,
	}

	page.Title = strings.TrimSpace(
		doc.Find("title").Text(),
	)

	doc.Find("a").Each(
		func(i int, s *goquery.Selection) {

			link, exists := s.Attr("href")

			if exists && link != "" {
				page.Links = append(
					page.Links,
					link,
				)
			}
		},
	)

	doc.Find("h1,h2,h3,h4,h5,h6").Each(
		func(i int, s *goquery.Selection) {

			heading := strings.TrimSpace(
				s.Text(),
			)

			if heading != "" {
				page.Headings = append(
					page.Headings,
					heading,
				)
			}
		},
	)

	return page
}

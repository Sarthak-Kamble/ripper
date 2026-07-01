package extractor

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Sarthak-Kamble/ripper/internal/model"
)

type LinkExtractor struct {
	baseURL string
}

func NewLinkExtractor(baseURL string) *LinkExtractor {
	return &LinkExtractor{
		baseURL: baseURL,
	}
}

func (e *LinkExtractor) Extract(doc *goquery.Document) []model.Link {

	var links []model.Link

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection){

		href, exists := s.Attr("href")

		if !exists {
			return
		}


		text := strings.TrimSpace(
			s.Text(),
		)


		linkType := e.getType(href)


		links = append(
			links,
			model.Link{
				URL: href,
				Text: text,
				Type: linkType,
			},
		)

	})


	return links

}

func (e *LinkExtractor) getType(link string) string {

	base, err := url.Parse(e.baseURL)

	if err != nil {
		return "unknown"
	}


	target, err := url.Parse(link)

	if err != nil {
		return "unknown"
	}


	if target.Host == "" {
		return "internal"
	}


	if target.Host == base.Host {
		return "internal"
	}


	return "external"
}
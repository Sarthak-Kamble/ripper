package commands

import (
	"fmt"

	"github.com/Sarthak-Kamble/ripper/internal/extractor"
	"github.com/Sarthak-Kamble/ripper/internal/scraper"
	"github.com/spf13/cobra"
)

var linksCmd = &cobra.Command{

	Use: "links [url]",

	Short: "Extract links from a webpage",

	Args: cobra.ExactArgs(1),


	RunE: func(cmd *cobra.Command, args []string) error {


		targetURL := args[0]


		fetcher := scraper.NewFetcher()


		html, _, err := fetcher.Fetch(targetURL)

		if err != nil {
			return err
		}



		parser := scraper.NewParser()


		doc, err := parser.Parse(html)

		if err != nil {
			return err
		}



		linkExtractor := extractor.NewLinkExtractor(
			targetURL,
		)



		links := linkExtractor.Extract(doc)



		for _, link := range links {

			fmt.Printf(
				"%s | %s | %s\n",
				link.Type,
				link.Text,
				link.URL,
			)

		}


		return nil
	},
}

func init(){
 
	
}
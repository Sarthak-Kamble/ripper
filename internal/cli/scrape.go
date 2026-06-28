package cli

import (
	"fmt"

	"github.com/Sarthak-Kamble/ripper/internal/scraper"
	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape [url]",
	Short: "Scrape a webpage",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		url := args[0]

		s := scraper.NewScraper()

		page, err := s.Scrape(url)

		if err != nil {
			fmt.Println("scrape failed:", err)
			return
		}

		fmt.Println("URL:")
		fmt.Println(page.URL)

		fmt.Println()

		fmt.Println("Status:")
		fmt.Println(page.Status)

		fmt.Println()

		fmt.Println("Title:")
		fmt.Println(page.Title)

		fmt.Println()

		fmt.Println("Links:")

		for _, link := range page.Links {
			fmt.Println("-", link)
		}

		fmt.Println()

		fmt.Println("Headings:")

		for _, heading := range page.Headings {
			fmt.Println("-", heading)
		}
	},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
}

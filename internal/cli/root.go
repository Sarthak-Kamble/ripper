package cli

import (
	"fmt"
	"os"

	"github.com/Sarthak-Kamble/ripper/internal/cli/commands"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "ripper",
	Short: "Ripper is a fast web scraping CLI.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}


func init() {

	commands.Register(
		rootCmd,
	)

}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "v1.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ripper version %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

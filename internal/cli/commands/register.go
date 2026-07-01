package commands

import "github.com/spf13/cobra"


func Register(root *cobra.Command) {

	root.AddCommand(
		linksCmd,
	)

}
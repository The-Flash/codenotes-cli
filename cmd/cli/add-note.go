package cli

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a note",
	Run:   func(cmd *cobra.Command, args []string) {},
}
